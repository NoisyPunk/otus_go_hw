package configs

import (
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/logger"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	LogLevel   string `yaml:"logLevel"`
	InmemStore bool   `yaml:"inmemStore"`
	Dsn        string `yaml:"postgres_dsn"`
}

func newConfig() *Config {
	return &Config{
		Host:       "",
		Port:       "",
		LogLevel:   "",
		InmemStore: false,
	}
}

func GetConfig(path string) *Config {
	logg := logger.New("debug")
	configYaml, err := os.ReadFile(path)
	if err != nil {
		logg.Fatal(err.Error())
	}
	config := newConfig()
	err = yaml.Unmarshal(configYaml, config)
	if err != nil {
		logg.Fatal(err.Error())
	}
	return config
}
