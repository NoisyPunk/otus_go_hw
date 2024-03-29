package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	LogLevel   string `yaml:"logLevel"`
	InmemStore bool   `yaml:"inmemStore"`
	Dsn        string `yaml:"postgresDsn"`
}

func newConfig() *Config {
	return &Config{
		Host:       "",
		Port:       "",
		LogLevel:   "",
		InmemStore: false,
	}
}

func GetConfig(path string) (*Config, error) {
	configYaml, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := newConfig()
	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
