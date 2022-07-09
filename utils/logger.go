package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	Logger  = NewLogger("common")
	logOnce sync.Once
)

func InitLogger(module string) {
	logOnce.Do(func() {
		Logger = NewLogger(module)
	})
}

func NewLogger(module string) *zap.Logger {
	stdoutLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.InfoLevel || level == zap.DebugLevel
	})
	stderrLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zap.WarnLevel || level == zapcore.ErrorLevel || level == zapcore.FatalLevel
	})
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	stdoutSyncer := zapcore.Lock(os.Stdout)
	stderrSyncer := zapcore.Lock(os.Stderr)
	core := zapcore.NewTee(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			stdoutSyncer,
			stdoutLevel,
		),
		zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			stderrSyncer,
			stderrLevel,
		),
	)
	logger := zap.New(core, zap.AddCaller())
	return logger.Named(module)
}
