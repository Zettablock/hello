package main

import (
	"github.com/Zettablock/zsource/utils"
)

// HandleBlock : shouldRetry tell zrunner whether to retry on errors
// *utils.Deps contains *gorm.DB which can be used for CRUD
// *utils.Deps also contains Logger
func HandleBlock(blockNumber string, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("HandleBlock1", "block number", blockNumber)
	//shouldRetry tells zrunner where to retry on errors
	shouldRetry := false
	return shouldRetry, nil
}

func HandleBlock2(blockNumber int, deps *utils.Deps) (bool, error) {
	deps.Logger.Info("HandleBlock2", "block number", blockNumber)
	//shouldRetry tells zrunner where to retry on errors
	shouldRetry := false
	return shouldRetry, nil
}
