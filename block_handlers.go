package main

import (
	"github.com/Zettablock/zsource/utils"
)

func HandleBlock(blockNumber string, deps *utils.Deps) (bool, error) {
	//shouldRetry tells zrunner where to retry on errors
	shouldRetry := false
	return shouldRetry, nil
}
