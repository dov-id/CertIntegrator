package indexer

import (
	"context"
	"math/big"
	"sort"
	"time"

	"github.com/dov-id/cert-integrator-svc/internal/data"
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

var logsHandlers = map[string]func(i *indexer, ctx context.Context, eventLog types.Log, client *ethclient.Client) error{
	crypto.Keccak256Hash([]byte(issuerTransferEventSignature)).Hex(): (*indexer).handleIssuerTransferLog,
	crypto.Keccak256Hash([]byte(fabricDeployEventSignature)).Hex():   (*indexer).handleFabricDeployLog,
}

func (i *indexer) Run(ctx context.Context) {
	go running.WithBackOff(
		ctx,
		i.log,
		serviceName,
		i.listen,
		data.IndexerTimeout*time.Second,
		data.IndexerTimeout*time.Second,
		data.IndexerTimeout*time.Second,
	)
}

func (i *indexer) listen(ctx context.Context) error {
	i.log.WithField("addresses", i.Addresses).Debugf("start listener")

	block, err := getBlockToStartFrom(i.ContractsQ, i.Addresses, i.Blocks)
	if err != nil {
		return errors.Wrap(err, "failed to get starting block")
	}

	err = i.processPastEvents(ctx, block, i.Clients[data.EthereumNetwork])
	if err != nil {
		return errors.Wrap(err, "failed to process past events")
	}

	err = i.subscribeAndProcessNewEvents(ctx, i.Clients[data.EthereumNetwork])
	if err != nil {
		return errors.Wrap(err, "failed to subscribe and process events:")
	}

	return nil
}

func getBlockToStartFrom(contractsQ data.Contracts, addresses []string, blocks []int64) (*big.Int, error) {
	dbContracts, err := contractsQ.FilterByAddresses(addresses...).Select()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contract")
	}

	for i := range dbContracts {
		blocks = append(blocks, dbContracts[i].Block)
	}

	sort.Slice(blocks, func(i, j int) bool { return blocks[i] > blocks[j] })

	// + 1 in order not to handle already handled last block
	return big.NewInt(blocks[0] + 1), nil
}

func (i *indexer) handleLogs(ctx context.Context, log types.Log, client *ethclient.Client) error {
	if logHandler, ok := logsHandlers[log.Topics[0].Hex()]; ok {
		err := logHandler(i, ctx, log, client)
		if err != nil {
			return errors.Wrap(err, "failed to handle log")
		}
	}

	return nil
}

func (i *indexer) processPastEvents(ctx context.Context, block *big.Int, client *ethclient.Client) error {
	i.log.WithFields(map[string]interface{}{"block": block.String(), "addresses": i.Addresses}).Debugf("start processing past events")

	filterQuery := ethereum.FilterQuery{
		Addresses: helpers.ConvertStringToAddresses(i.Addresses),
		FromBlock: block,
		ToBlock:   nil,
	}

	oldLogs, err := client.FilterLogs(ctx, filterQuery)
	if err != nil {
		return errors.Wrap(err, "failed to filter logs")
	}

	for _, log := range oldLogs {
		i.log.WithFields(map[string]interface{}{"block": log.BlockNumber, "address": log.Address.Hex()}).Debugf("processing past event")

		err = i.handleLogs(ctx, log, client)
		if err != nil {
			return errors.Wrap(err, "failed to handle log")
		}

		block = big.NewInt(int64(log.BlockNumber))
	}

	i.log.WithFields(map[string]interface{}{"block": block.String(), "addresses": i.Addresses}).Debugf("finish processing past events")
	return nil
}

func (i *indexer) subscribeAndProcessNewEvents(ctx context.Context, client *ethclient.Client) error {
	query := ethereum.FilterQuery{
		Addresses: helpers.ConvertStringToAddresses(i.Addresses),
	}

	logs := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to logs")
	}

	for {
		select {
		case err = <-subscription.Err():
			return errors.Wrap(err, "some error with subscription")
		case vLog := <-logs:
			if err = i.handleLogs(ctx, vLog, client); err != nil {
				return errors.Wrap(err, "failed to handle log")
			}
		case address := <-i.issuerCh:
			i.Addresses = append(i.Addresses, address)
			subscription.Unsubscribe()
			return i.listen(ctx)
		}
	}
}
