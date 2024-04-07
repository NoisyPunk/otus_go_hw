package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	EventServerPort string `yaml:"eventServerPort"`
	LogLevel        string `yaml:"logLevel"`
	InmemStore      bool   `yaml:"inmemStore"`
	Dsn             string `yaml:"postgresDsn"`
}

func GetConfig(path string) (*Config, error) {
	configYaml, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
