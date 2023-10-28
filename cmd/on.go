package cmd

import (
	"github.com/go-shecan/operations"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(turnOn)
}

var turnOn = &cobra.Command{
	Use:   "on",
	Short: "this will change resolv.conf file to servers in config file",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		operations.ReadSaveDeleteDefaultServers_WriteNewServers()
	},
}
