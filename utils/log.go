package utils

import (
	"sync"

	"go.uber.org/zap"
	)

var logger *zap.SugaredLogger
var once sync.Once

func NewLogger() *zap.SugaredLogger {
	once.Do(func() {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		cfg.OutputPaths = []string{
			"/tmp/gostatus.log",
		}
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		l, err :=  cfg.Build()
		if err != nil {
			panic(err)
		}
		logger = l.Sugar()
	})
	return logger
}