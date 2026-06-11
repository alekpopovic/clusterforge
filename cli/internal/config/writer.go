package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Write(path string, cfg *Config, force bool) error {
	if !force {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("%s already exists; use --force to overwrite it", path)
		}
	}
	cfg.ApplyDefaults()
	if err := cfg.Validate(); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("encode config: %w", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write config %s: %w", path, err)
	}
	return nil
}

func (c *Config) Save(path string) error {
	return Write(path, c, true)
}
