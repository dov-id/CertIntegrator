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

const (
	serviceName                  = "indexer"
	IssuerContract               = "CertIssuer"
	FabricContract               = "CertFabric"
	ZeroAddress                  = "0x0000000000000000000000000000000000000000"
	issuerTransferEventSignature = "Transfer(address,address,uint256)"
	fabricDeployEventSignature   = "TokenContractDeployed(address,(uint256,string,string))"
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

	client, err := ethclient.Dial(i.cfg.Infura().Sepolia + i.cfg.Infura().Key)
	if err != nil {
		return errors.Wrap(err, "failed to make dial connect")
	}

	block, err := getBlockToStartFrom(i.ContractsQ, i.Addresses, i.Blocks)
	if err != nil {
		return errors.Wrap(err, "failed to get starting block")
	}

	err = i.processPastEvents(block, client)
	if err != nil {
		return errors.Wrap(err, "failed to process past events")
	}

	err = i.subscribeAndProcessNewEvents(client)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe and process events:")
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

	sort.Slice(blocks, func(i, j int) bool { return blocks[i] > blocks[j] })

	return big.NewInt(blocks[0]), nil
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

func (i *Indexer) processPastEvents(block *big.Int, client *ethclient.Client) error {
	i.log.WithFields(map[string]interface{}{"block": block.String(), "addresses": i.Addresses}).Debugf("start processing past events")

	filterQuery := ethereum.FilterQuery{
		Addresses: helpers.ConvertStringToAddresses(i.Addresses),
		FromBlock: block,
		ToBlock:   nil,
	}

	oldLogs, err := client.FilterLogs(context.Background(), filterQuery)
	if err != nil {
		return errors.Wrap(err, "failed to filter logs")
	}

	for _, log := range oldLogs {
		i.log.WithFields(map[string]interface{}{"block": log.BlockNumber, "address": log.Address.Hex()}).Debugf("processing past event")

		err = i.handleLogs(log, client)
		if err != nil {
			return errors.Wrap(err, "failed to handle log")
		}

		block = big.NewInt(int64(log.BlockNumber))
	}

	i.log.WithFields(map[string]interface{}{"block": block.String(), "addresses": i.Addresses}).Debugf("finish processing past events")
	return nil
}

func (i *Indexer) subscribeAndProcessNewEvents(client *ethclient.Client) error {
	query := ethereum.FilterQuery{
		Addresses: helpers.ConvertStringToAddresses(i.Addresses),
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
			if err = i.handleLogs(vLog, client); err != nil {
				return errors.Wrap(err, "failed to handle log")
			}
		}
	}
}
