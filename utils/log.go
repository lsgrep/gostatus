package utils

import (
	"sync"

	"github.com/spf13/viper"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var once sync.Once

func init() {
	viper.SetDefault("log", "/tmp/gostatus.log")
}

func NewLogger() *zap.SugaredLogger {
	once.Do(func() {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		logFile := viper.GetString("log")
		cfg.OutputPaths = []string{logFile}
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		l, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		logger = l.Sugar()
	})
	return logger
}
