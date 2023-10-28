package db

import (
	"git.mills.io/prologic/bitcask"
)

func Connect() (*bitcask.Bitcask, error) {
	db, error := bitcask.Open("")
	if error != nil {
		return nil, error
	}
	return db, nil
}

func SaveResolvFileState(db *bitcask.Bitcask, nameServers []byte) {
	db.Put([]byte(""), nameServers)
}

func GetResolvFileLastState(db *bitcask.Bitcask) ([]byte, error) {
	allKeys := db.Keys()
	for key := range allKeys {
		if string(key) == "" {
			return db.Get(key)
		}
	}
	return nil, nil
}
