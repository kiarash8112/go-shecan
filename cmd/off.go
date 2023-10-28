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
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		operations.RestoreResolv()
	},
}
