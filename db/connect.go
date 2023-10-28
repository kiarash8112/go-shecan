package db

import (
	"git.mills.io/prologic/bitcask"
)

func Connect() (*bitcask.Bitcask, error) {
	db, error := bitcask.Open("/tmp/db")
	if error != nil {
		return nil, error
	}
	return db, nil
}

func SaveResolvFileState(db *bitcask.Bitcask, nameServers []byte) {
	db.Put([]byte("resolvFile"), nameServers)
}

func GetResolvFileLastState(db *bitcask.Bitcask) ([]byte, error) {
	allKeys := db.Keys()
	for key := range allKeys {
		if string(key) == "resolvFile" {
			return db.Get(key)
		}
	}
	return nil, nil
}
