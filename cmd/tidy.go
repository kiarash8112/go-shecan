package cmd

import (
	"fmt"
	"os"

	"github.com/go-shecan/operations"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(modTidy)
}

var modTidy = &cobra.Command{
	Use:   "tidy",
	Short: "this command will change dns servers and run go mod tidy",
	Run: func(cmd *cobra.Command, args []string) {
		out, err := operations.RunTidy()
		fmt.Printf("command-output: %v\n", out)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(2)
		}
	},
}
