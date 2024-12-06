package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	Venv         string `yaml:"venv"`
	Parser       string `yaml:"parser"`
}

var appcof Config

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", ".\\config\\local.yaml")
		flag.Parse()

		configPath = *flags
		if configPath == "" {
			log.Fatal("missing config file path")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	err := cleanenv.ReadConfig(configPath, &appcof)
	if err != nil {
		log.Fatalf("cannot read config: %v", err)
	}

	return &appcof
}
