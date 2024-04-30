package main

import (
	"github.com/Zettablock/zsource/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slog"
	"os"
	"testing"
)

func TestEventHandlers(t *testing.T) {
	t.Run("test example", func(t *testing.T) {
		jsonHandler := slog.NewJSONHandler(os.Stdout, nil)
		logger := slog.New(jsonHandler)
		shouldRetry, err := HandleBlock("123", &utils.Deps{
			Logger: logger,
		})
		assert.NoError(t, err)
		assert.False(t, shouldRetry)
	})
}
