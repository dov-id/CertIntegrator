package storage

import (
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
)

func AddAttempt(storage DailyStorage, address common.Address, maxAttemptsAmount int64) error {
	hexAddress := address.Hex()
	var newValue int64 = 1

	val, ok := storage.Get(hexAddress)
	if !ok {
		if maxAttemptsAmount < newValue {
			return data.ErrMaxAttemptsAmount
		}
		storage.Set(hexAddress, newValue)
		return nil
	}

	attempts, ok := val.(int64)
	if !ok {
		return data.ErrFailedToCastInt
	}
	newValue = attempts + 1

	if maxAttemptsAmount < newValue {
		return data.ErrMaxAttemptsAmount
	}

	storage.Set(hexAddress, newValue)

	return nil
}

func GetRemainingAttempts(storage DailyStorage, address common.Address) (int64, error) {
	val, ok := storage.Get(address.Hex())
	if !ok {
		return 0, data.ErrNoSuchKey
	}

	attempts, ok := val.(int64)
	if !ok {
		return 0, data.ErrFailedToCastInt
	}

	return attempts, nil
}

func CheckAttemptsExceeded(storage DailyStorage, address common.Address, maxAttemptsAmount int64) (bool, error) {
	val, ok := storage.Get(address.Hex())
	if !ok {
		return false, nil
	}

	attempts, ok := val.(int64)
	if !ok {
		return false, data.ErrFailedToCastInt
	}

	return maxAttemptsAmount == attempts, nil
}
