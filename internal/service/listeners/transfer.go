package listeners

import (
	"fmt"

	"github.com/dov-id/CertIntegrator/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

type LogTransfer struct {
	//event Transfer(address indexed from, address indexed to, uint value);
	From common.Address
	To   common.Address
}

func (l *Listener) handleIssuerTransferLog(eventLog types.Log) error {
	var event LogTransfer

	event.From = common.HexToAddress(eventLog.Topics[1].Hex())
	event.To = common.HexToAddress(eventLog.Topics[2].Hex())

	if event.From.Hex() == "0x0000000000000000000000000000000000000000" {
		//handle mint
		fmt.Printf("From: %s\n", event.From.Hex())
		fmt.Printf("To: %s\n", event.To.Hex())

		err := l.BlocksQ.Upsert(data.Block{
			ContractName:    issuerContract,
			LastBlockNumber: int64(eventLog.BlockNumber),
		})
		if err != nil {
			return errors.Wrap(err, "failed to save last handled block")
		}

		return nil
	}

	//handle transfer
	fmt.Printf("From: %s\n", event.From.Hex())
	fmt.Printf("To: %s\n", event.To.Hex())

	err := l.BlocksQ.Upsert(data.Block{
		ContractName:    issuerContract,
		LastBlockNumber: int64(eventLog.BlockNumber),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	return nil
}
