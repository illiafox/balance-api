package config

import (
	"fmt"
	"path"

	"github.com/ilyakaznacheev/cleanenv"
)

func New(file string) (Config, error) {
	file = path.Clean(file)

	var cfg Config

	err := cleanenv.ReadConfig(file, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}

	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("read environment: %w", err)
	}

	return cfg, nil
}
