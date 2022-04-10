package logger

import (
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

type loggerEnvironment struct {
	filePath string
	stdout   bool
	level    zapcore.Level
}

func getEnv() (*loggerEnvironment, error) {
	res := loggerEnvironment{}
	res.filePath = os.Getenv("LOGGER_FILE_PATH")

	var err error
	res.stdout, err = strconv.ParseBool(os.Getenv("LOGGER_STDOUT"))
	if err != nil {
		res.stdout = true
	}

	level := os.Getenv("LOGGER_LEVEL")
	switch level {
	case "debug":
		res.level = zapcore.DebugLevel
	case "error":
		res.level = zapcore.ErrorLevel
	default:
		res.level = zapcore.InfoLevel
	}

	return &res, nil
}
