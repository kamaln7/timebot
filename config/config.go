package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

// Config is a struct containing the different config
// options for sonarrhook
type Config struct {
	Host      string
	InChannel bool
	Timezones map[string]string
}

// Read reads the config file and returns a Config struct
func Read() (*Config, error) {
	configfile := "./config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		return nil, fmt.Errorf("config file is missing: path = %s: %v", configfile, err)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		return nil, fmt.Errorf("invalid config: %v", err)
	}

	return &config, nil
}
