package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-shecan",
	Short: "go-shecan is dns changer for downloading packages in macos and linux",
	Long:  `go-shecan change user resolv.conf file to use config file dns server(shecan.ir by default) `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Missing command! Try 'go-shecan --help' for more information.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
