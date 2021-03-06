package xlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Color      bool
	LogFile    string
	JsonLog    bool
	Caller     bool
	StackTrace bool
	Level      string // debug, info, warn, error
}

func New(cfg Config) *zap.Logger {
	var zConfig zap.Config

	zConfig.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	zConfig.DisableCaller = !cfg.Caller
	zConfig.DisableStacktrace = !cfg.StackTrace

	var level zapcore.Level

	if cfg.Level == "" {
		cfg.Level = "debug"
	}

	err := level.Set(cfg.Level)
	if err != nil {
		fmt.Println("invalid zap log level: ", cfg.Level)
		return nil
	}

	zConfig.Level = zap.NewAtomicLevelAt(level)
	if cfg.JsonLog {
		zConfig.Encoding = "json"
	} else {
		zConfig.Encoding = "console"
	}
	if cfg.Color {
		zConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if cfg.LogFile == "" {
		zConfig.OutputPaths = []string{"stdout"}
		zConfig.ErrorOutputPaths = []string{"stderr"}
	} else {
		zConfig.OutputPaths = []string{cfg.LogFile}
		zConfig.ErrorOutputPaths = []string{cfg.LogFile}
	}

	option := []zap.Option{
		zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel),
	}

	logger, err := zConfig.Build(option...)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return logger
}
