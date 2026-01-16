package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Load(env, logPath, logFile string) *zap.Logger {
	var config zap.Config

	if env == "prod" {
		config = zap.NewProductionConfig()
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		config.Encoding = "json"
	} else {
		config = zap.NewDevelopmentConfig()
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		config.Encoding = "console"
	}

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	stdout := zapcore.Lock(os.Stdout)
	stderr := zapcore.Lock(os.Stderr)

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%v/%v", logPath, logFile),
		MaxSize:    50, // 50MB
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	})

	consoleEncoder := zapcore.NewConsoleEncoder(config.EncoderConfig)
	jsonEncoder := zapcore.NewJSONEncoder(config.EncoderConfig)

	var streamEncoder zapcore.Encoder

	if config.Encoding == "console" {
		streamEncoder = consoleEncoder
	} else {
		streamEncoder = jsonEncoder
	}

	core := zapcore.NewTee(
		zapcore.NewCore(streamEncoder, stdout, config.Level),
		zapcore.NewCore(jsonEncoder, fileWriter, config.Level),
		zapcore.NewCore(streamEncoder, stderr, zap.ErrorLevel),
	)

	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Fields(
			zap.Int("pid", os.Getpid()),
			zap.String("env", env),
		),
	)
	zap.ReplaceGlobals(logger)

	return logger
}
