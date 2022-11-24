package config

import (
	"flag"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	host = "127.0.0.1"
	port = "8080"

	dbURL    = "host=localhost dbname=dev sslmode=disabled"
	dbDriver = "postgres"

	loggerLevel = "debug"

	certPath = ""
	keyPath  = ""

	endpoint        = "0.0.0.0/cloud"
	assessKeyID     = "admin"
	secretAccessKey = "secretkey"
	useSSL          = false

	allowCredentials = true

	tokenLength = 32
	contextKey  = "csrf"
	contextName = "csrf_token"
	csrfMaxAge  = 86400

	policy        = "{\"Version\": \"2012-10-17\",\"Statement\": [{\"Action\": [\"s3:GetObject\"],\"Effect\": \"Allow\",\"Principal\": {\"AWS\": [\"*\"]},\"Resource\": [\"arn:aws:s3:::$(bucket)/*\"],\"Sid\": \"\"}]}"
	expire        = 60
	symbolsToHash = 1
)

var (
	allowMethods = []string{
		http.MethodDelete,
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
	}
	allowHeaders = []string{
		"Content-Type",
		"Content-length",
	}
	allowOrigins = []string{
		"https://vdonate.ml",
		"http://localhost:8080",
		"http://localhost:4200",
	}
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

	CORS struct {
		AllowMethods     []string `yaml:"allow_methods"`
		AllowHeaders     []string `yaml:"allow_headers"`
		AllowCredentials bool     `yaml:"allow_credentials"`
		AllowOrigins     []string `yaml:"allow_origins"`
	} `yaml:"cors"`

	CSRF struct {
		Status      bool   `yaml:"status"`
		TokenLength uint8  `yaml:"token_length"`
		ContextKey  string `yaml:"context_key"`
		ContextName string `yaml:"context_name"`
		MaxAge      int    `yaml:"max_age"`
	} `yaml:"csrf"`

	S3 struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		UseSSL          bool   `yaml:"use_ssl"`
		Buckets         struct {
			Policy        string `yaml:"policy"`
			SymbolsToHash int    `yaml:"symbols_to_hash"`
			Expire        int    `yaml:"expire"`
		} `yaml:"buckets"`
	} `yaml:"s3"`
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
		}{
			Host:     host,
			Port:     port,
			CertPath: certPath,
			KeyPath:  keyPath,
		}),

		DB: struct {
			Driver string `yaml:"driver"`
			URL    string `yaml:"url"`
		}(struct {
			Driver string
			URL    string
		}{
			Driver: dbDriver,
			URL:    dbURL,
		}),

		Logger: struct {
			Level string `yaml:"level"`
		}{
			Level: loggerLevel,
		},

		Deploy: struct {
			Mode bool `yaml:"mode"`
		}{
			Mode: false,
		},

		CORS: struct {
			AllowMethods     []string `yaml:"allow_methods"`
			AllowHeaders     []string `yaml:"allow_headers"`
			AllowCredentials bool     `yaml:"allow_credentials"`
			AllowOrigins     []string `yaml:"allow_origins"`
		}{
			AllowMethods:     allowMethods,
			AllowHeaders:     allowHeaders,
			AllowCredentials: allowCredentials,
			AllowOrigins:     allowOrigins,
		},

		CSRF: struct {
			Status      bool   `yaml:"status"`
			TokenLength uint8  `yaml:"token_length"`
			ContextKey  string `yaml:"context_key"`
			ContextName string `yaml:"context_name"`
			MaxAge      int    `yaml:"max_age"`
		}{
			Status:      false,
			TokenLength: tokenLength,
			ContextKey:  contextKey,
			ContextName: contextName,
			MaxAge:      csrfMaxAge,
		},

		S3: struct {
			Endpoint        string `yaml:"endpoint"`
			AccessKeyID     string `yaml:"access_key_id"`
			SecretAccessKey string `yaml:"secret_access_key"`
			UseSSL          bool   `yaml:"use_ssl"`
			Buckets         struct {
				Policy        string `yaml:"policy"`
				SymbolsToHash int    `yaml:"symbols_to_hash"`
				Expire        int    `yaml:"expire"`
			} `yaml:"buckets"`
		}{
			Endpoint:        endpoint,
			AccessKeyID:     assessKeyID,
			SecretAccessKey: secretAccessKey,
			UseSSL:          useSSL,
			Buckets: struct {
				Policy        string `yaml:"policy"`
				SymbolsToHash int    `yaml:"symbols_to_hash"`
				Expire        int    `yaml:"expire"`
			}(struct {
				Policy        string
				SymbolsToHash int
				Expire        int
			}{Policy: policy, SymbolsToHash: symbolsToHash, Expire: expire}),
		},
	}
}

func (c *Config) Open(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// Start YAML decoding from file
	if err = yaml.NewDecoder(file).Decode(&c); err != nil {
		return err
	}

	return nil
}

func PathFlag(path *string) {
	flag.StringVar(path, "config-path", "./configs/config_local.yaml", "path to config file")
}
