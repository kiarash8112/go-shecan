package main

import (
	"fmt"
	"io/ioutil"
)

func mai2n() {
	f, err := ioutil.ReadFile("/etc/resolv.conf")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(string(f))
}
