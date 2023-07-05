package helpers

import (
	"github.com/dov-id/cert-integrator-svc/internal/config"
	"github.com/dov-id/cert-integrator-svc/internal/data"
	"github.com/ethereum/go-ethereum/common"
)

func SeparateContractArrays(list []config.Contract) ([]int64, []string) {
	blocks := make([]int64, 0)
	addresses := make([]string, 0)

	for i := range list {
		blocks = append(blocks, list[i].FromBlock)
		addresses = append(addresses, list[i].Address)
	}

	return blocks, addresses
}

func ConvertStringToAddresses(addrs []string) []common.Address {
	addresses := make([]common.Address, 0)

	for i := range addrs {
		addresses = append(addresses, common.HexToAddress(addrs[i]))
	}

	return addresses
}

func SeparateDataContractArrays(list []data.Contract) ([]int64, []string) {
	blocks := make([]int64, 0)
	addresses := make([]string, 0)

	for i := range list {
		blocks = append(blocks, list[i].Block)
		addresses = append(addresses, list[i].Address)
	}

	return blocks, addresses
}
