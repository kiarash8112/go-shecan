package operations

import (
	"github.com/go-shecan/db"
	"github.com/go-shecan/resolv"
)

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

	err = resolv.WriteDnsServersToResolvFile(file)
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

	err = resolv.ResetResolvFileToLastState(file, string(content))
	if err != nil {
		return err
	}

	return nil
}
