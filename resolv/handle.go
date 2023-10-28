package resolv

import (
	"fmt"
	"io"
	"os"
)

func GetResolvFile() (*os.File, error) {
	resolvFile, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_RDONLY, os.ModeAppend)
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
	if err := os.Truncate("/etc/resolv.conf", 0); err != nil {
		return err
	}
	return nil
}

func WriteShecanServersToResolvFile(file *os.File) error {
	shecanServers := `nameserver 185.51.200.2
nameserver 178.22.122.100`
	_, err := file.WriteString(shecanServers)
	if err != nil {
		return err
	}
	return nil
}
