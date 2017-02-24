package config

import (
	"log"
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
func Read() Config {
	configfile := "./config.toml"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}
