package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-shecan/db"
)

func main() {

	readFile, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	totalRes := []string{}

	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), "nameserver") {
			res := strings.Split(fileScanner.Text(), "nameserver")
			totalRes = append(totalRes, res[1][1:])
		}
	}

	db := db.Connect()
	for _, obj := range totalRes {
		db.Put([]byte(obj), nil)
	}

	allKeys := db.Keys()
	for key := range allKeys {
		fmt.Printf("key: %v\n", string(key))
	}

	if err := os.Truncate("/etc/resolv.conf", 0); err != nil {
		log.Printf("Failed to truncate: %v", err)
	}

	_, err = readFile.WriteString("nameserver 185.51.200.2")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
