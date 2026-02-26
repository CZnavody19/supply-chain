package setup

import (
	"os"

	"github.com/CZnavody19/supply-chain/src/config"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger(configuration config.Config) {
	cores := []zapcore.Core{}
	if configuration.LoggingConfig.EnableDebugLogger {
		lowCore, highCore := setupDebugLogger()
		cores = append(cores, lowCore)
		cores = append(cores, highCore)
	}
	if configuration.LoggingConfig.EnableFileLogger {
		cores = append(cores, setupFileLogger(configuration.LoggingConfig))
	}
	core := zapcore.NewTee(cores...)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
}

func setupFileLogger(logConfig config.LoggingConfig) zapcore.Core {
	var level zapcore.Level
	err := level.Set(logConfig.FileLogLevel)
	if err != nil {
		// We want to panic here, since logger is important
		panic(err)
	}

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	if logConfig.FileLogOutput == "stdout" {
		sink, _, err := zap.Open(logConfig.FileLogOutput)
		if err != nil {
			// We want to panic here, since logger is important
			panic(err)
		}
		return ecszap.NewCore(encoderConfig, sink, level)
	} else {
		// rolling file appender
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConfig.FileLogOutput,
			MaxSize:    logConfig.FileMaxSize,
			MaxBackups: logConfig.FileMaxBackups,
			MaxAge:     logConfig.FileMaxAge,
		})

		return ecszap.NewCore(encoderConfig, w, level)
	}
}

func setupDebugLogger() (zapcore.Core, zapcore.Core) {
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > zap.ErrorLevel
	})

	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zap.ErrorLevel
	})

	consoleWriter := zapcore.Lock(os.Stdout)
	consoleError := zapcore.Lock(os.Stderr)

	return zapcore.NewCore(consoleEncoder, consoleWriter, priority), zapcore.NewCore(consoleEncoder, consoleError, errorPriority)
}
