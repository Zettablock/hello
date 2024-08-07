package main

import (
	"fmt"
	"time"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

func HandleTransfer(log ethereum.Log, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("HandleTransfer", "block number", log.BlockNumber, "pipeline_name", deps.Config.GetPipelineName())
	fmt.Println("HandleTransfer called in proj3")
	fmt.Println("log", log)
	fmt.Println("block number", log.BlockNumber)
	fmt.Println("log index", log.LogIndex)
	fmt.Println("log block time", log.BlockTime)
	fmt.Println("log block date", log.BlockDate)
	fmt.Println("log event", log.Event)
	time.Sleep(2 * time.Second)
	return false, nil
}
