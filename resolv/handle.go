package resolv

import (
	"fmt"
	"io"
	"os"
)

func GetResolvFile() (*os.File, error) {
	resolvFile, err := os.OpenFile("", os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	return resolvFile, nil
}

func ReadResolvFile(resolvFile *os.File) ([]byte, error) {
	nameServers, err := io.ReadAll(resolvFile)
	if err != nil {
		return nil, err
	}
	return nameServers, nil
}

func ClearResolvFile() error {
	if err := os.Truncate("", 0); err != nil {
		return err
	}
	return nil
}

func WriteDnsServersToResolvFile(file *os.File) error {
	dnsServers := ""
	_, err := file.WriteString(dnsServers)
	if err != nil {
		return err
	}
	return nil
}
