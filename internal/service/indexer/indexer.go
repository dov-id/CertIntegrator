package indexer

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/data/postgres"
	"github.com/dov-id/cert-integrator-svc/internal/helpers"
	"github.com/ethereum/go-ethereum"
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
	IssuerContract               = "CertIssuer"
	FabricContract               = "CertFabric"
	ZeroAddress                  = "0x0000000000000000000000000000000000000000"
	issuerTransferEventSignature = "Transfer(address,address,uint256)"
	fabricDeployEventSignature   = "Deploy(string,address)"
)

var logsHandlers = map[string]func(i *Indexer, eventLog types.Log, client *ethclient.Client) error{
	crypto.Keccak256Hash([]byte(issuerTransferEventSignature)).Hex(): (*Indexer).handleIssuerTransferLog,
	crypto.Keccak256Hash([]byte(fabricDeployEventSignature)).Hex():   (*Indexer).handleFabricDeployLog,
}

func NewIndexer(cfg config.Config, ctx context.Context, addresses []string, blocks []int64, cancel context.CancelFunc) IIndexer {
	return &Indexer{
		cfg:        cfg,
		ctx:        ctx,
		log:        cfg.Log(),
		Addresses:  addresses,
		Blocks:     blocks,
		ContractsQ: postgres.NewContractsQ(cfg.DB().Clone()),
		Cancel:     cancel,
	}
}

func (i *Indexer) Run(ctx context.Context) {
	go running.WithBackOff(
		ctx,
		i.log,
		serviceName,
		i.listen,
		30*time.Second,
		30*time.Second,
		30*time.Second,
	)
}

func (i *Indexer) listen(_ context.Context) error {
	i.log.WithField("addresses", i.Addresses).Debugf("start listener")

	client, err := ethclient.Dial(i.cfg.Infura().Link + i.cfg.Infura().Key)
	if err != nil {
		return errors.Wrap(err, "failed to make dial connect")
	}

	block, err := getBlockToStartFrom(i.ContractsQ, i.Addresses, i.Blocks)
	if err != nil {
		return errors.Wrap(err, "failed to get starting block")
	}

	query := ethereum.FilterQuery{
		Addresses: helpers.ConvertStringToAddresses(i.Addresses),
		FromBlock: block,
	}

	logs := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(i.ctx, query, logs)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to logs")
	}

	for {
		select {
		case err = <-subscription.Err():
			return errors.Wrap(err, "some error with subscription")
		case vLog := <-logs:
			if err = i.handleLogs(vLog, client); err != nil {
				return errors.Wrap(err, "failed to handle log")
			}
		}
	}

}

func (i *Indexer) handleLogs(log types.Log, client *ethclient.Client) error {
	if logHandler, ok := logsHandlers[log.Topics[0].Hex()]; ok {
		err := logHandler(i, log, client)
		if err != nil {
			return err
		}
	}

	return nil
}

func getBlockToStartFrom(contractsQ data.Contracts, addresses []string, blocks []int64) (*big.Int, error) {
	contracts, err := contractsQ.FilterByAddresses(addresses...).Select()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract")
	}

	for i := range contracts {
		blocks = append(blocks, contracts[i].Block)
	}

	sort.Slice(blocks, func(i, j int) bool { return blocks[i] < blocks[j] })

	return big.NewInt(blocks[0]), nil
}
