package db

import (
	"git.mills.io/prologic/bitcask"
)

var DBInstance DB

func Set(d DB) {
	DBInstance = d
}
func Get() DB {
	return DBInstance
}

type DB struct {
	db *bitcask.Bitcask
}

func Connect(dbPath string) (DB, error) {
	db, error := bitcask.Open(dbPath)
	if error != nil {
		return DB{}, error
	}
	return DB{db: db}, nil
}

func (d DB) SaveResolvFileState(nameServers []byte) {
	d.db.Put([]byte(""), nameServers)
}

func (d DB) GetResolvFileLastState() ([]byte, error) {
	allKeys := d.db.Keys()
	for key := range allKeys {
		if string(key) == "" {
			return d.db.Get(key)
		}
	}
	return nil, nil
}
