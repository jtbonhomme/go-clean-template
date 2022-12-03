package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		Log    `yaml:"logger"`
		Server `yaml:"server"`
	}

	// App -.
	App struct {
		Name string `env-required:"true" yaml:"name"    env:"APP_NAME"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"logLevel" env:"LOG_LEVEL"`
	}

	// Server -.
	Server struct {
		Port string `env-required:"true" yaml:"port" env:"PORT"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
