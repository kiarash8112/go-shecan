package db

import (
	"git.mills.io/prologic/bitcask"
	config "github.com/go-shecan/configs"
)

var DBInstance DB

func Set(d DB) {
	DBInstance = d
}
func Get() DB {
	return DBInstance
}

func Connect(c config.Config) (DB, error) {
	db, error := bitcask.Open(c.DBPath)
	if error != nil {
		return DB{}, error
	}
	return DB{db: db, storeKey: c.StoreKey}, nil
}
