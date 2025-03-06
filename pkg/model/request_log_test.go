package model

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestLogInfo(t *testing.T) {
	var buff bytes.Buffer
	var errorBuff bytes.Buffer
	logger := zap.New(
		zapcore.NewCore(zapcore.NewJSONEncoder(
			zapcore.EncoderConfig{}), zapcore.AddSync(&buff),
			zapcore.DebugLevel,
		),
		zap.ErrorOutput(zapcore.AddSync(&errorBuff)),
	)

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	r := RequestLog{
		RequestAddress:    "http://localhost:8080",
		StatusCode:        200,
		Method:            "GET",
		ResponseTimeMs:    100,
		TimeoutTimeMs:     1000,
		SpeedTestResultID: 1,
	}

	r.LogInfo()

	stringBuffer := buff.String()
	require.Contains(t, stringBuffer, "content")
}
