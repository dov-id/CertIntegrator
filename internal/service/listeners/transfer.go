package listeners

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogTransfer struct {
	//event Transfer(address indexed from, address indexed to, uint value);
	From common.Address
	To   common.Address
}

func handleIssuerTransferLog(eventLog types.Log) error {
	var event LogTransfer

	event.From = common.HexToAddress(eventLog.Topics[1].Hex())
	event.To = common.HexToAddress(eventLog.Topics[2].Hex())

	if event.From.Hex() == "0x0000000000000000000000000000000000000000" {
		//handle mint
		fmt.Printf("From: %s\n", event.From.Hex())
		fmt.Printf("To: %s\n", event.To.Hex())

		//TODO: save this in database
		fmt.Printf("Log Block Number: %d\n", eventLog.BlockNumber)
		return nil
	}

	//handle transfer
	fmt.Printf("From: %s\n", event.From.Hex())
	fmt.Printf("To: %s\n", event.To.Hex())

	//TODO: save this in database
	fmt.Printf("Log Block Number: %d\n", eventLog.BlockNumber)

	return nil
}
