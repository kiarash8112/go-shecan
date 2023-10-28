package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-shecan",
	Short: "go-shecan is dns changer for downloading packages in macos and linux",
	Long: `go-shecan change user resolv.conf file to use shecan.ir dns servers and after getting 
	packages it will restore resolv.conf to first state
	you can also turn go-shecan on to do some work and after that turn it off manually`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello-World")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
