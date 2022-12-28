package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config struct to store all app config.
	Config struct {
		App  `yaml:"app"`
		Db   `yaml:"db"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}

	// App struct to store all app information.
	App struct {
		Env     string `env-required:"true" yaml:"env"     env:"APP_ENV" env-description:"App env to display some helper (dev,test,prod)"`
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME" env-description:"Application name"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION" env-description:"Application version"`
	}

	// Db struct to store all database information.
	Db struct {
		Dsn string `env-required:"true" yaml:"dsn" env:"DB_DSN" env-description:"Database Domaine Server Name that the application must be linked"`
	}

	// HTTP struct to store all http server information.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT" env-description:"Port that the server will listen to handle http requests"`
	}

	// Log struct to store all log information.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL" env-description:"Log level. trace < debug < info < warn < error < fatal < panic"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// default config load by file depending on where we call goapp
	// In godog test it will be called in tests folder and then will load /src/tests/config/config.yml file
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
