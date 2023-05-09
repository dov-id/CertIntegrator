package listeners

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogDeploy struct {
	//event Deploy(string, address indexed contract);
	CourseName    string
	CourseAddress common.Address
}

func handleFabricDeployLog(eventLog types.Log) error {
	var event LogDeploy

	event.CourseName = eventLog.Topics[1].String()
	event.CourseAddress = common.HexToAddress(eventLog.Topics[2].Hex())

	//handle deploy course
	fmt.Printf("Name: %s\n", event.CourseName)
	fmt.Printf("Addr: %s\n", event.CourseAddress.Hex())

	//TODO: save this in database
	fmt.Printf("Log Block Number: %d\n", eventLog.BlockNumber)

	return nil
}
