package main

import (
	"github.com/lsgrep/gostatus/bar"
	"github.com/lsgrep/gostatus/utils"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var logger = utils.NewLogger()

func main() {
	logger.Debug("gostatus has been started")
	pflag.String("config", "config.yml", "config file")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	configFile := viper.GetString("config")

	statusBar := bar.NewGoStatusBar()
	statusBar.Run(configFile)
}
