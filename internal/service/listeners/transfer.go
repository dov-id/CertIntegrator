package listeners

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	sql "github.com/iden3/go-merkletree-sql/db/pgx/v2"
	"github.com/iden3/go-merkletree-sql/v2"
	"github.com/pkg/errors"
)

type LogTransfer struct {
	//event Transfer(address indexed from, address indexed to, uint value);
	From common.Address
	To   common.Address
}

func (l *Listener) handleIssuerTransferLog(eventLog types.Log) error {
	l.log.Infof("start handling transfer event")

	var event LogTransfer

	event.From = common.HexToAddress(eventLog.Topics[1].Hex())
	event.To = common.HexToAddress(eventLog.Topics[2].Hex())

	contract, err := l.AddressesQ.FilterByAddresses(l.Address.Hex()).Get()
	if err != nil {
		return errors.Wrap(err, "failed to get address from database")
	}

	if contract == nil {
		contract, err = l.AddressesQ.Insert(data.Contract{
			Name:    IssuerContract,
			Address: l.Address.Hex(),
		})
		if err != nil {
			return errors.Wrap(err, "failed to save new contract")
		}
	}

	treeStorage := sql.NewSqlStorage(l.DBConn, contract.Id)

	mTree, err := merkletree.NewMerkleTree(l.ctx, treeStorage, 100)
	if err != nil {
		return errors.Wrap(err, "failed to get merkle tree")
	}

	//handle mint
	if event.From.Hex() == ZeroAddress {
		err = mTree.Add(l.ctx, event.To.Big(), event.To.Big())
		if err != nil {
			return errors.Wrap(err, "failed to add new leaf in merkle tree")
		}

		err = l.BlocksQ.Upsert(data.Block{
			ContractAddress: l.Address.Hex(),
			LastBlockNumber: int64(eventLog.BlockNumber),
		})
		if err != nil {
			return errors.Wrap(err, "failed to save last handled block")
		}

		l.log.Infof("finish handling mint event")
		return nil
	}

	//handle transfer
	err = mTree.Delete(l.ctx, event.To.Big())
	if err != nil {
		return errors.Wrap(err, "failed to delete address from merkle tree")
	}

	err = l.BlocksQ.Upsert(data.Block{
		ContractAddress: l.Address.Hex(),
		LastBlockNumber: int64(eventLog.BlockNumber),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save last handled block")
	}

	l.log.Infof("finish handling transfer event")
	return nil
}
