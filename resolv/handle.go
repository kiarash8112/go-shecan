package resolv

import (
	"fmt"
	"io"
	"os"
	"reflect"

	config "github.com/go-shecan/configs"
)

var resolvInstance Resolve

func Set(r Resolve) {
	resolvInstance = r
}
func Get() Resolve {
	return resolvInstance
}

type Resolve struct {
	*config.DnsServers
	resolvPath string
}

func (r Resolve) GetResolvFile() (*os.File, error) {
	resolvFile, err := os.OpenFile(r.resolvPath, os.O_APPEND|os.O_RDWR|os.O_CREATE|os.O_RDONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

	return resolvFile, nil
}

func (r Resolve) ReadResolvFile(resolvFile *os.File) ([]byte, error) {
	nameServers, err := io.ReadAll(resolvFile)
	if err != nil {
		return nil, err
	}
	return nameServers, nil
}

func (r Resolve) ResetResolvFileToLastState(file *os.File, lastContents string) error {
	_, err := file.WriteString(lastContents)
	if err != nil {
		return err
	}
	return nil
}

func (r Resolve) ClearResolvFile() error {
	if err := os.Truncate(r.resolvPath, 0); err != nil {
		return err
	}
	return nil
}

func (r Resolve) WriteDnsServersToResolvFile(file *os.File) error {
	v := reflect.ValueOf(r.DnsServers)

	for i := 0; i < v.NumField(); i++ {
		_, err := file.WriteString(v.Field(i).String())
		if err != nil {
			return err
		}
	}

	return nil
}
