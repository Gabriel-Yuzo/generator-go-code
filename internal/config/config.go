package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
}

var AppConfig *Config

func InitConfig() {
	configFile, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	AppConfig = &Config{}
	err = decoder.Decode(AppConfig)
	if err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}
}
