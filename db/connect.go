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

type DB struct {
	storeKey string
	db       *bitcask.Bitcask
}

func Connect(c config.Config) (DB, error) {
	db, error := bitcask.Open(c.DBPath)
	if error != nil {
		return DB{}, error
	}
	return DB{db: db, storeKey: c.StoreKey}, nil
}

func (d DB) SaveResolvFileState(nameServers []byte) {
	d.db.Put([]byte(d.storeKey), nameServers)
}

func (d DB) GetResolvFileLastState() ([]byte, error) {
	allKeys := d.db.Keys()
	for key := range allKeys {
		if string(key) == d.storeKey {
			return d.db.Get(key)
		}
	}
	return nil, nil
}
