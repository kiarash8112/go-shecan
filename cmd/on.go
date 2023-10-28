package cmd

import (
	"fmt"
	"os"

	"github.com/go-shecan/operations"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(turnOn)
}

var turnOn = &cobra.Command{
	Use:   "on",
	Short: "this will change resolv.conf file to servers in config file",
	Run: func(cmd *cobra.Command, args []string) {
		err := operations.ReadSaveDeleteDefaultServers_WriteNewServers()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(2)
		}
	},
}
