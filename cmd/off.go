package cmd

import (
	"fmt"
	"os"

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
		err := operations.RestoreResolv()
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(2)
		}
	},
}
