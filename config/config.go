package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	TLD string `yaml:"tld"`
}

func readConfig() *Config {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	config := Config{}
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
	return &config
}

var DefaultConfig *Config = readConfig()
