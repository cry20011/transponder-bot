package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	Token string `env:"transponder_telegram_bot_token" envDefault:"token"`
}

func New() (*Config, error) {
	var c Config

	if err := env.Parse(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
