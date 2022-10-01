package storage_config

type Config struct {
	DatabaseURL string `toml:"database_url"`
}

func New() *Config {
	return &Config{}
}
