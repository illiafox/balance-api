package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func New() (Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("read environment: %w", err)
	}

	return cfg, nil
}
