package cmd

import (
	"github.com/go-shecan/operations"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(turnOff)
}

var turnOff = &cobra.Command{
	Use:   "off",
	Short: "this will restore resolv.conf to default state",
	Run: func(cmd *cobra.Command, args []string) {
		operations.RestoreResolv()
	},
}
