package main

import (
	"fmt"
	"time"

	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
)

// HandleBlock : shouldRetry tell zrunner whether to retry on errors
// *utils.Deps contains *gorm.DB which can be used for CRUD
// *utils.Deps also contains Logger
func HandleBlock(block ethereum.Block, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("HandleBlock", "block number", block.Number, "pipeline_name", deps.Config.GetPipelineName())
	fmt.Println("HandleBlock called in proj3")
	fmt.Println("block", block)
	fmt.Println("block number", block.Number)
	fmt.Println("block hash", block.Hash)
	fmt.Println("block timestamp", block.Timestamp)
	fmt.Println("block date", block.BlockDate)
	time.Sleep(10 * time.Second)
	return false, nil
}
