package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates package specific logging pipeline.
func New(name string) *zap.SugaredLogger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	level := zap.InfoLevel
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level)
	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	return logger.Named(name).Sugar()
}
