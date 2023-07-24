package helpers

import (
	"github.com/ethereum/go-ethereum/common"
)

func ConvertStringToAddresses(addrs []string) []common.Address {
	addresses := make([]common.Address, 0)

	for i := range addrs {
		addresses = append(addresses, common.HexToAddress(addrs[i]))
	}

	return addresses
}
