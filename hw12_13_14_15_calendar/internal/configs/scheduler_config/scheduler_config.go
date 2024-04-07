package schedulerconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host                   string `yaml:"host"`
	Port                   string `yaml:"port"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	LogLevel               string `yaml:"logLevel"`
	Dsn                    string `yaml:"postgresDsn"`
	RemoveScannerFrequency int    `yaml:"removeScannerFrequency"`
	NotifyScannerFrequency int    `yaml:"notifyScannerFrequency"`
	StoragePeriod          int    `yaml:"storagePeriod"`
}

func newConfig() *Config {
	return &Config{}
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
