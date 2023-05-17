package listeners

import (
	"context"
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
	"github.com/jackc/pgx/v4"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/distributed_lab/running"
)

//CertIssuer - ERC-721 events:
//		Mint (Transfer from 0x0 address to someone), which contains information about for whom NFT is minted
//		Transfer, which contains information about from who and to whom NFT is transferred
//CertFabric - common fabric contract events:
//		Deploy, which contains information about the course name and its address

const (
	serviceName    = "listener"
	IssuerContract = "CertIssuer"
	FabricContract = "CertFabric"
	ZeroAddress    = "0x0000000000000000000000000000000000000000"
	//infuraLink                   = "wss://mainnet.infura.io/ws/v3/"
	infuraLink                   = "wss://goerli.infura.io/ws/v3/"
	issuerTransferEventSignature = "Transfer(address,address,uint256)"
	fabricDeployEventSignature   = "Deploy(string,address)"
)

var logsHandlers = map[string]func(l *Listener, eventLog types.Log) error{
	crypto.Keccak256Hash([]byte(issuerTransferEventSignature)).Hex(): (*Listener).handleIssuerTransferLog,
	crypto.Keccak256Hash([]byte(fabricDeployEventSignature)).Hex():   (*Listener).handleFabricDeployLog,
}

type IListener interface {
	Run(ctx context.Context)
}

type Listener struct {
	cfg config.Config
	ctx context.Context
	log *logan.Entry

	DBConn     *pgx.Conn
	Address    common.Address
	FromBlock  *big.Int
	BlocksQ    data.Blocks
	AddressesQ data.Contracts
}

func NewListener(cfg config.Config, ctx context.Context, conn *pgx.Conn, address string, fromBlock int64) IListener {
	return &Listener{
		cfg:        cfg,
		ctx:        ctx,
		log:        cfg.Log().WithField("address", address),
		DBConn:     conn,
		Address:    common.HexToAddress(address),
		FromBlock:  big.NewInt(fromBlock),
		BlocksQ:    postgres.NewBlocksQ(cfg.DB().Clone()),
		AddressesQ: postgres.NewContractsQ(cfg.DB().Clone()),
	}
}

func (l *Listener) Run(ctx context.Context) {
	go running.WithBackOff(
		ctx,
		l.log,
		serviceName,
		l.listen,
		30*time.Second,
		30*time.Second,
		30*time.Second,
	)
}

func (l *Listener) listen(_ context.Context) error {
	l.log.Infof("start listener")

	client, err := ethclient.Dial(l.cfg.Infura().Link + l.cfg.Infura().Key)
	if err != nil {
		return errors.Wrap(err, "failed to make dial connect")
	}

	block, err := l.BlocksQ.FilterByContractAddress(l.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get block")
	}
	if block != nil {
		l.FromBlock = big.NewInt(block.LastBlockNumber)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{l.Address},
		FromBlock: l.FromBlock,
	}

	logs := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(l.ctx, query, logs)
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
			return err
		}
	}

	return nil
}
