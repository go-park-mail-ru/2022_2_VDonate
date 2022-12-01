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

	dbURL    = "host=localhost dbname=dev sslmode=disable"
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

	Services struct {
		Auth struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"auth"`
		Donates struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"donates"`
		Images struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"images"`
		Posts struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"posts"`
		Subscribers struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"subscribers"`
		Subscriptions struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"subscriptions"`
		Users struct {
			Port        string `yaml:"port"`
			MetricsPort string `yaml:"metricsPort"`
			Host        string `yaml:"host"`
		} `yaml:"users"`
	} `yaml:"services"`

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
		Enabled     bool   `yaml:"enabled"`
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

		Services: struct {
			Auth struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"auth"`
			Donates struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"donates"`
			Images struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"images"`
			Posts struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"posts"`
			Subscribers struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"subscribers"`
			Subscriptions struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"subscriptions"`
			Users struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			} `yaml:"users"`
		}(struct {
			Auth struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
			Donates struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
			Images struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
			Posts struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
			Subscribers struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
			Subscriptions struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
			Users struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}
		}{
			Auth: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8081",
				MetricsPort: "9081",
				Host:        "0.0.0.0",
			}),
			Donates: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8082",
				MetricsPort: "9082",
				Host:        "0.0.0.0",
			}),
			Images: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8083",
				MetricsPort: "9083",
				Host:        "0.0.0.0",
			}),
			Posts: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8084",
				MetricsPort: "9084",
				Host:        "0.0.0.0",
			}),
			Subscribers: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8085",
				MetricsPort: "9085",
				Host:        "0.0.0.0",
			}),
			Subscriptions: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8086",
				MetricsPort: "9086",
				Host:        "0.0.0.0",
			}),
			Users: struct {
				Port        string `yaml:"port"`
				MetricsPort string `yaml:"metricsPort"`
				Host        string `yaml:"host"`
			}(struct {
				Port        string
				MetricsPort string
				Host        string
			}{
				Port:        "8087",
				MetricsPort: "9087",
				Host:        "0.0.0.0",
			}),
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
			Enabled     bool   `yaml:"enabled"`
			TokenLength uint8  `yaml:"token_length"`
			ContextKey  string `yaml:"context_key"`
			ContextName string `yaml:"context_name"`
			MaxAge      int    `yaml:"max_age"`
		}{
			Enabled:     false,
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
	flag.StringVar(path, "config-path", "./cmd/api/configs/config_local.yaml", "path to config file")
}
