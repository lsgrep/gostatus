package log

import (
	"sync"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var once sync.Once

func init() {
	cfg := zap.NewDevelopmentConfig()
	l, e := cfg.Build()
	if e != nil {
		panic(e)
	}
	logger = l.Sugar()
}

func ConfigureLogger(file string) {
	once.Do(func() {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig = zap.NewDevelopmentEncoderConfig()
		cfg.OutputPaths = []string{file}
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		l, err := cfg.Build()
		if err != nil {
			panic(err)
		}
		logger = l.Sugar()
	})
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}