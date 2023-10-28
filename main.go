package main

import (
	"fmt"
	"runtime"

	"github.com/go-shecan/cmd"
	i "github.com/go-shecan/init"
)

func init() {
	i.Init()
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("can't execute this program on windows machine")
	} else {
		cmd.Execute()
	}
}
