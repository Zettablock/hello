package main

import (
	"errors"

	"github.com/Zettablock/zsource/utils"
)

// HandleBlock : shouldRetry tell zrunner whether to retry on errors
// *utils.Deps contains *gorm.DB which can be used for CRUD
// *utils.Deps also contains Logger
func HandleBlock(blockNumber string, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("HandleBlock1", "block number", blockNumber, "pipeline_name2", deps.Config.GetPipelineName())
	////shouldRetry tells zrunner where to retry on errors
	shouldRetry := false
	return shouldRetry, nil
}

func HandleBlock2(blockNumber int, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("HandleBlock2", "block number", blockNumber)
	return true, errors.New("block handler fake error to test")
}
