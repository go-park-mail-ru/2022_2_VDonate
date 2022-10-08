package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	host         = "127.0.0.1"
	port         = "8080"
	dbURL        = "host=localhost dbname=dev sslmode=disabled"
	dbDriver     = "postgres"
	requestDebug = true
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	DB struct {
		Driver string `yaml:"driver"`
		URL    string `yaml:"url"`
	} `yaml:"db"`

	Debug struct {
		Request bool `yaml:"request"`
	} `yaml:"debug"`
}

func New() *Config {
	return &Config{
		Server: struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		}(struct {
			Host string
			Port string
		}{Host: host, Port: port}),

		DB: struct {
			Driver string `yaml:"driver"`
			URL    string `yaml:"url"`
		}(struct {
			Driver string
			URL    string
		}{Driver: dbDriver, URL: dbURL}),

		Debug: struct {
			Request bool `yaml:"request"`
		}{Request: requestDebug},
	}
}

func (c *Config) Open(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// Start YAML decoding from file
	if err := yaml.NewDecoder(file).Decode(&c); err != nil {
		return err
	}

	return nil
}

func PathFlag(path *string) {
	flag.StringVar(path, "config-path", "./configs/config_local.yaml", "path to config file")
}
