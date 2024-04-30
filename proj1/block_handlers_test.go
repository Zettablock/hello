package main

import (
	"github.com/Zettablock/zsource/dao/ethereum"
	"github.com/Zettablock/zsource/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
	"os"
	"testing"
)

func TestBlockHandlers(t *testing.T) {
	t.Run("test example", func(t *testing.T) {
		jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
		logger := slog.New(jsonHandler)
		shouldRetry, err := HandleAVSMetadataURIUpdated(ethereum.Log{}, &utils.Deps{
			Logger: logger,
		})
		assert.Error(t, err)
		assert.True(t, shouldRetry)
	})
}
