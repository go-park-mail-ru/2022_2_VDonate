package config

import (
	"flag"
	storage_config "github.com/go-park-mail-ru/2022_2_VDonate/internal/storages/config"
)

const (
	port   = ":8080"
	dbName = "postgres"
)

type Config struct {
	Port    string `toml:"port"`
	DbName  string `toml:"dbName"`
	Storage *storage_config.Config
}

func New() *Config {
	return &Config{
		Port:    port,
		DbName:  dbName,
		Storage: storage_config.New(),
	}
}

func PathFlag(path *string) {
	flag.StringVar(path, "config-path", "./configs/config.toml", "path to config file")
}
