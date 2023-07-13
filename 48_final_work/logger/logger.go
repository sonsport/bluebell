package logger

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init(mode string) {
	encoder := getEncoder()
	writeSync := getWriteSync()
	var core zapcore.Core
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSync, zapcore.DebugLevel),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
		fmt.Println("dev")
	} else {
		core = zapcore.NewCore(encoder, writeSync, zapcore.DebugLevel)
	}
	// 这里为什么不能用 settings.ViperConfig.Log.Level
	logger = zap.New(core)
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "useTime"
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriteSync() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
