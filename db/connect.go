package db

import (
	"git.mills.io/prologic/bitcask"
)

type DB struct {
	storeKey string
	db       *bitcask.Bitcask
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
