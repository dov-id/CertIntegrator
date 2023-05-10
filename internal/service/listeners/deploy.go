package listeners

import (
	"github.com/dov-id/CertIntegrator/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type LogDeploy struct {
	//event Deploy(string, address indexed contract);
	CourseName    string
	CourseAddress common.Address
}

func (l *Listener) handleFabricDeployLog(eventLog types.Log) error {
	var event LogDeploy

	event.CourseName = eventLog.Topics[1].String()
	event.CourseAddress = common.HexToAddress(eventLog.Topics[2].Hex())

	err := l.AddressesQ.Insert(data.Address{
		CourseName: event.CourseName,
		Address:    event.CourseAddress.Hex(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save new contract")
	}

	err = l.BlocksQ.Upsert(data.Block{
		ContractName:    fabricContract,
		LastBlockNumber: int64(eventLog.BlockNumber),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	return nil
}
