package helpers

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/dov-id/cert-integrator-svc/internal/service/storage"
	"github.com/ethereum/go-ethereum/common"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Add(storage storage.DailyStorage, address common.Address, maxAttemptsAmount int64) error {
	hexAddress := address.Hex()
	var newValue int64 = 1

	val, ok := storage.Get(hexAddress)
	if !ok {
		if maxAttemptsAmount < newValue {
			return errors.New(data.MaxAttemptsAmountErr)
		}
		storage.Set(hexAddress, newValue)
		return nil
	}

	attempts, ok := val.(int64)
	if !ok {
		return errors.New(data.FailedToCastIntErr)
	}
	newValue = attempts + 1

	if maxAttemptsAmount < newValue {
		return errors.New(data.MaxAttemptsAmountErr)
	}

	storage.Set(hexAddress, newValue)

	return nil
}

func Get(storage storage.DailyStorage, address common.Address) (int64, error) {
	val, ok := storage.Get(address.Hex())
	if !ok {
		return 0, errors.New(data.NoSuchKeyErr)
	}

	attempts, ok := val.(int64)
	if !ok {
		return 0, errors.New(data.FailedToCastIntErr)
	}

	return attempts, nil
}

func Check(storage storage.DailyStorage, address common.Address, maxAttemptsAmount int64) (bool, error) {
	val, ok := storage.Get(address.Hex())
	if !ok {
		return false, nil
	}

	attempts, ok := val.(int64)
	if !ok {
		return false, errors.New(data.FailedToCastIntErr)
	}

	return maxAttemptsAmount == attempts, nil
}
