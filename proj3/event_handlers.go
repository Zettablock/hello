package main

import (
	"fmt"
	"time"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

func HandleTransfer(log ethereum.Log, deps *utils.Deps) (bool, error) {
	fmt.Println("HandleTransfer called in proj3")
	fmt.Println("log", log)
	fmt.Println("block number", log.BlockNumber)
	fmt.Println("log index", log.LogIndex)
	time.Sleep(10 * time.Second)
	return false, nil
}
