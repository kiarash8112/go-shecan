package db

import (
	"git.mills.io/prologic/bitcask"
)

func Connect() *bitcask.Bitcask {
	db, _ := bitcask.Open("/tmp/db")
	return db
}
