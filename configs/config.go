package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type DnsServers struct {
	Server1 string `json:"server1"`
	Server2 string `json:"server2"`
}
type Config struct {
	DBPath     string `json:"db_path"`
	DnsServers `json:"dns_servers"`
	StoreKey   string `json:"store_key"`
	ResolvPath string `json:"resolv_path"`
}

func LoadAndConvert(path string) (Config, error) {
	d, err := Load(path)
	if err != nil {
		return Config{}, err
	}
	return Convert(d)
}

func Load(path string) ([]byte, error) {
	return os.ReadFile(filepath.Clean(path))
}

func Convert(d []byte) (Config, error) {
	c := Config{}
	if err := json.Unmarshal(d, &c); err != nil {
		return Config{}, err
	}
	return c, nil
}
