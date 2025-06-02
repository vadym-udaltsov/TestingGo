package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type EnvConfig struct {
	BaseURL string `yaml:"baseUrl"`
	Token   string `yaml:"token"`
}

var Cfg EnvConfig

func LoadConfig() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	type RawConfig struct {
		Default string    `yaml:"default"`
		Dev     EnvConfig `yaml:"dev"`
		Staging EnvConfig `yaml:"staging"`
	}

	file, err := os.ReadFile("config/env.yaml")
	if err != nil {
		log.Fatalf("Error during config reading: %v", err)
	}

	var raw RawConfig
	if err := yaml.Unmarshal(file, &raw); err != nil {
		log.Fatalf("YAML parsing error: %v", err)
	}

	switch env {
	case "staging":
		Cfg = raw.Staging
	default:
		Cfg = raw.Dev
	}
}
