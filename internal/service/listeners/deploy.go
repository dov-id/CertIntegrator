package listeners

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	sql "github.com/iden3/go-merkletree-sql/db/pgx/v2"
	"github.com/iden3/go-merkletree-sql/v2"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type LogDeploy struct {
	//event Deploy(string, address indexed contract);
	CourseName    string
	CourseAddress common.Address
}

func (l *Listener) handleFabricDeployLog(eventLog types.Log) error {
	l.log.Infof("start handling deploy event")

	var event LogDeploy

	event.CourseName = eventLog.Topics[1].String()
	event.CourseAddress = common.HexToAddress(eventLog.Topics[2].Hex())

	newContract, err := l.AddressesQ.Insert(data.Contract{
		Name:    event.CourseName,
		Address: event.CourseAddress.Hex(),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save new contract")
	}

	treeStorage := sql.NewSqlStorage(l.DBConn, newContract.Id)

	_, err = merkletree.NewMerkleTree(l.ctx, treeStorage, 100)
	if err != nil {
		return errors.Wrap(err, "failed to create new merkle tree")
	}

	NewListener(
		l.cfg,
		l.ctx,
		l.DBConn,
		newContract.Address,
		int64(eventLog.BlockNumber),
	).Run(l.ctx)

	err = l.BlocksQ.Upsert(data.Block{
		ContractAddress: l.Address.Hex(),
		LastBlockNumber: int64(eventLog.BlockNumber),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	l.log.Infof("finish handling deploy event")
	return nil
}
