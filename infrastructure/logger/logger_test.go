package logger_test

import (
	"apex_roulette/application_service/service"
	"apex_roulette/infrastructure/logger"
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructLog(t *testing.T) {
	t.Run("構造体のログが出力できること", func(t *testing.T) {
		var buf bytes.Buffer
		logger := logger.NewLogger(&buf)
		logger.StructLog(service.LogLevelInfo, "test")

		type log struct {
			Serverity string `json:"severity"`
			Message   string `json:"message"`
		}
		var actual log
		err := json.Unmarshal(buf.Bytes(), &actual)
		assert.NoError(t, err)
		expected := log{
			Serverity: "Info",
			Message:   "test",
		}
		assert.Equal(t, expected, actual)
	})
}
