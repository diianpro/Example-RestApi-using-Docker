package config

import (
	"context"
	"github.com/caarlos0/env/v6"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Port     int
	DBName   string
	Password string
	Username string
	Host     string
	*DbConfig
}

func NewConfig(ctx context.Context) (*Config, error) {
	cfg := Config{
		DbConfig: &DbConfig{},
	}
	if err := env.Parse(&cfg); err != nil {
		log.Errorf("NewConfig: %v", err)
	}
	err := cfg.Init(ctx)
	if err != nil {
		log.Errorf("error in config: initialize store config: %v", err)
		return nil, err
	}

	return &cfg, nil
}
