package nono

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func ConfigureLogger() *zap.SugaredLogger {
	config := zap.NewDevelopmentConfig()
	config.DisableCaller = true
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()

	return logger.Sugar()
}
