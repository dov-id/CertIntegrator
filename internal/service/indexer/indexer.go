package indexer

import (
	"context"
	"math/big"

	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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
		i.cfg.Timeouts().Indexer,
		i.cfg.Timeouts().Indexer,
		i.cfg.Timeouts().Indexer,
	)
}

func (i *indexer) listen(ctx context.Context) error {
	i.log.WithField("addresses", i.Addresses).Debugf("start listener")

	err := i.processPastEvents(ctx, i.Clients[data.EthereumNetwork])
	if err != nil {
		return errors.Wrap(err, "failed to process past events")
	}

	err = i.subscribeAndProcessNewEvents(ctx, i.Clients[data.EthereumNetwork])
	if err != nil {
		return errors.Wrap(err, "failed to subscribe and process events:")
	}

	return nil
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

func (i *indexer) processPastEvents(ctx context.Context, client *ethclient.Client) error {
	i.log.WithField("addresses", i.Addresses).Debugf("start processing past events")

	for k := 0; k < len(i.Addresses); k++ {
		filterQuery := ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(i.Addresses[k])},
			FromBlock: big.NewInt(i.Blocks[i.Addresses[k]] + 1),
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

			i.Blocks[log.Address.Hex()] = int64(log.BlockNumber)
		}
	}

	i.log.WithField("addresses", i.Addresses).Debugf("finish processing past events")
	return nil
}

func (i *indexer) subscribeAndProcessNewEvents(ctx context.Context, client *ethclient.Client) error {
	i.log.WithField("addresses", i.Addresses).Debugf("subscribing to new events")

	addresses, err := convertStringsToAddresses(i.Addresses)
	if err != nil {
		return errors.Wrap(err, "failed to convert strings to addresses")
	}

	logs := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(
		ctx,
		ethereum.FilterQuery{
			Addresses: addresses,
		},
		logs,
	)
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

func convertStringsToAddresses(addrs []string) ([]common.Address, error) {
	addresses := make([]common.Address, 0)

	for i := range addrs {
		if !common.IsHexAddress(addrs[i]) {
			return nil, data.ErrInvalidEthAddress
		}
		addresses = append(addresses, common.HexToAddress(addrs[i]))
	}

	return addresses, nil
}
