package operations

import (
	"log"

	"github.com/go-shecan/db"
	"github.com/go-shecan/resolv"
)

var headerSize int = 35

func ReadSaveDeleteDefaultServers_WriteNewServers() error {
	db := db.Get()
	resolv := resolv.Get()

	file, err := resolv.GetResolvFile()
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := resolv.ReadResolvFile(file)
	if err != nil {
		return err
	}

	db.SaveResolvFileState(content)

	err = resolv.ClearResolvFile()
	if err != nil {
		return err
	}

	dnsServers, err := resolv.WriteDnsServersToResolvFile(file)
	log.Println("server changed from", string(content[headerSize:]), "to", dnsServers)

	if err != nil {
		return err
	}

	return nil
}

func RestoreResolv() error {
	db := db.Get()
	resolv := resolv.Get()

	content, err := db.GetResolvFileLastState()
	if err != nil {
		return err
	}

	file, err := resolv.GetResolvFile()
	if err != nil {
		return err
	}

	err = resolv.ClearResolvFile()
	if err != nil {
		return err
	}

	err = resolv.ResetResolvFileToLastState(file, string(content))
	if err != nil {
		return err
	}
	log.Println("changed to", string(content[headerSize:]))
	return nil
}
