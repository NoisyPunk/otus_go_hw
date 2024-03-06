package configs

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	ErrReadConfigFile = fmt.Errorf("can't read config file")
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	LogLevel string `yaml:"level"`
}

func newConfig() *Config {
	return &Config{
		Host:     "",
		Port:     "",
		LogLevel: "",
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
