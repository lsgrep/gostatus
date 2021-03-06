package main

import (
	"github.com/lsgrep/gostatus/bar"
	"github.com/lsgrep/gostatus/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	pflag.String("config", "gostatus.yml", "config file")
	pflag.String("log", "/tmp/gostatus.log", "log file")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}

	l := viper.GetString("log")
	log.ConfigureLogger(l)
	log.Debug("gostatus has been started")
	configFile := viper.GetString("config")
	statusBar := bar.NewGoStatusBar()
	statusBar.Run(configFile)
}
