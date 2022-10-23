package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	host        = "127.0.0.1"
	port        = "8080"
	dbURL       = "host=localhost dbname=dev sslmode=disabled"
	dbDriver    = "postgres"
	loggerLevel = "debug"
	certPath    = ""
	keyPath     = ""
)

type Config struct {
	Server struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		CertPath string `yaml:"cert_path"`
		KeyPath  string `yaml:"key_path"`
	} `yaml:"server"`

	DB struct {
		Driver string `yaml:"driver"`
		URL    string `yaml:"url"`
	} `yaml:"db"`

	Logger struct {
		Level string `yaml:"level"`
	} `yaml:"debug"`

	Deploy struct {
		Mode bool `yaml:"mode"`
	} `yaml:"deploy"`
}

func New() *Config {
	return &Config{
		Server: struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			CertPath string `yaml:"cert_path"`
			KeyPath  string `yaml:"key_path"`
		}(struct {
			Host     string
			Port     string
			CertPath string
			KeyPath  string
		}{Host: host, Port: port, CertPath: certPath, KeyPath: keyPath}),

		DB: struct {
			Driver string `yaml:"driver"`
			URL    string `yaml:"url"`
		}(struct {
			Driver string
			URL    string
		}{Driver: dbDriver, URL: dbURL}),

		Logger: struct {
			Level string `yaml:"level"`
		}{Level: loggerLevel},

		Deploy: struct {
			Mode bool `yaml:"mode"`
		}{Mode: false},
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
