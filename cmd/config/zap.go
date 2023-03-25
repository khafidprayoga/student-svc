package config

import (
	"go.uber.org/zap"
)

func GetZapLogger() *zap.Logger {
	// TODO: separate dev/prod logging config
	logger, _ := zap.NewDevelopment()
	_ = logger.Sync()
	return logger
}
