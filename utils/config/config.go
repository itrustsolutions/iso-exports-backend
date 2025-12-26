package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	App      AppConfig
	Server   ServerConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Env  string `env:"ITRUST_APP_ENV" envDefault:"development"`
	Name string `env:"ITRUST_APP_NAME" envDefault:"iTrustApp"`
}

type ServerConfig struct {
	Port               string   `env:"ITRUST_PORT,required"`
	ReadTimeout        int      `env:"ITRUST_READ_TIMEOUT" envDefault:"15"`
	WriteTimeout       int      `env:"ITRUST_WRITE_TIMEOUT" envDefault:"15"`
	IdleTimeout        int      `env:"ITRUST_IDLE_TIMEOUT" envDefault:"15"`
	CORSAllowedOrigins []string `env:"ITRUST_CORS_ALLOWED_ORIGINS,required"`
}

type DatabaseConfig struct {
	Host            string `env:"ITRUST_DB_HOST,required"`
	Port            int    `env:"ITRUST_DB_PORT,required"`
	User            string `env:"ITRUST_DB_USER,required"`
	Password        string `env:"ITRUST_DB_PASSWORD,required"`
	Name            string `env:"ITRUST_DB_NAME,required"`
	SSLMode         string `env:"ITRUST_DB_SSLMODE" envDefault:"disable"`
	MaxOpenConns    int    `env:"ITRUST_DB_MAX_OPEN_CONNS" envDefault:"25"`
	MaxIdleConns    int    `env:"ITRUST_DB_MAX_IDLE_CONNS" envDefault:"25"`
	ConnMaxLifetime int    `env:"ITRUST_DB_CONN_MAX_LIFETIME" envDefault:"5"`
	ConnMaxIdleTime int    `env:"ITRUST_DB_CONN_MAX_IDLE_TIME" envDefault:"5"`
}

var (
	instance *Config
	once     sync.Once
)

// Load the config once and return the same instance on subsequent calls & exit the process on error with status 1
func GetConfigOrExist() *Config {
	once.Do(func() {
		cfg := &Config{}
		if err := env.Parse(cfg); err != nil {
			fmt.Fprintln(os.Stderr, "Could not parse environment variables:", err)
			os.Exit(1)
		}
		instance = cfg
	})
	return instance
}
