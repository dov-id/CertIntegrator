package listeners

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/dov-id/CertIntegrator/internal/config"
	"github.com/dov-id/CertIntegrator/internal/data"
	"github.com/dov-id/CertIntegrator/internal/data/postgres"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

//CertIssuer - ERC-721 events:
//		Mint (Transfer from 0x0 address to someone), which contains information about for whom NFT is minted
//		Transfer, which contains information about from who and to whom NFT is transferred
//CertFabric - common fabric contract events:
//		Deploy, which contains information about the course name and its address

const (
	serviceName                  = "listener"
	issuerContract               = "CertIssuer"
	fabricContract               = "CertFabric"
	infuraLink                   = "wss://mainnet.infura.io/ws/v3/"
	issuerTransferEventSignature = "Transfer(address,address,uint256)"
	fabricDeployEventSignature   = "Deploy(string,address)"
)

var logsHandlers = map[string]func(l *Listener, eventLog types.Log) error{
	crypto.Keccak256Hash([]byte(issuerTransferEventSignature)).Hex(): (*Listener).handleIssuerTransferLog,
	crypto.Keccak256Hash([]byte(fabricDeployEventSignature)).Hex():   (*Listener).handleFabricDeployLog,
}

type IListener interface {
	Run(ctx context.Context, cfg config.Config)
}

type Listener struct {
	Address    common.Address
	InfuraKey  string
	FromBlock  *big.Int
	BlocksQ    data.Blocks
	AddressesQ data.ContractAddresses
}

func NewListener(cfg config.Config, address string, fromBlock int64) IListener {
	var blockToStart *big.Int = nil

	return &Listener{
		Address:    common.HexToAddress(address),
		InfuraKey:  cfg.Infura().Key,
		FromBlock:  blockToStart,
		BlocksQ:    postgres.NewBlocksQ(cfg.DB().Clone()),
		AddressesQ: postgres.NewContractAddressesQ(cfg.DB().Clone()),
	}
}

func (l *Listener) Run(ctx context.Context, cfg config.Config) {
	go running.WithBackOff(
		ctx,
		cfg.Log(),
		serviceName,
		l.listen,
		30*time.Second,
		30*time.Second,
		30*time.Second,
	)
}

func (l *Listener) listen(_ context.Context) error {
	fmt.Println("start listeners")
	client, err := ethclient.Dial(infuraLink + l.InfuraKey)
	if err != nil {
		return errors.Wrap(err, "failed to make dial connect")
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.Address},
		FromBlock: l.FromBlock,
	}

	logs := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to logs")
	}

	for {
		select {
		case err = <-subscription.Err():
			return errors.Wrap(err, "some error with subscription")
		case vLog := <-logs:
			if err = l.handleLogs(vLog); err != nil {
				return errors.Wrap(err, "failed to handle log")
			}
		}
	}

}

func (l *Listener) handleLogs(log types.Log) error {
	if logHandler, ok := logsHandlers[log.Topics[0].Hex()]; ok {
		err := logHandler(l, log)
		if err != nil {
			return errors.Wrap(err, "failed to handle log")
		}
	}

	return nil
}
