package config

import "sync"

var (
	once     sync.Once
	instance *Config
)

type Config struct {
	JWTKey string
	DNS    string
}

// GetConfig devuelve la instancia de Config.
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			JWTKey: "u';Q1I07RyE2@",
			DNS:    "geovany:123456@tcp(127.0.0.1:3306)/go",
		}
	})

	return instance
}
