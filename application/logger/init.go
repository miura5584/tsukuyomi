package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func init() {
	envVals, err := getEnv()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var outputPaths []string
	if envVals.filePath != "" {
		outputPaths = append(outputPaths, envVals.filePath)
	}

	if envVals.stdout {
		outputPaths = append(outputPaths, "stdout")
	}

	logConfig := zap.Config{
		OutputPaths: outputPaths,
		Level:       zap.NewAtomicLevelAt(envVals.level),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			CallerKey:    "caller",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	if Log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}
