package init

import (
	"log"
	"path/filepath"

	config "github.com/go-shecan/configs"
	"github.com/go-shecan/db"
	"github.com/go-shecan/resolv"
)

func Init() {
	path, err := filepath.Abs("configs/config.json")
	if err != nil {
		log.Fatal("faild to load configs : ", err)
	}

	c, err := config.LoadAndConvert(path)
	if err != nil {
		log.Fatal("faild to load configs : ", err)
	}

	d, err := db.Connect(c)
	if err != nil {
		panic("can't connect to database")
	}
	db.Set(d)

	r := resolv.New(c)
	resolv.Set(r)
}
