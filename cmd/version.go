package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print current version of chain daemon",
	Long:  `print current version of chain daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}
