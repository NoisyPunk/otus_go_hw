package schedulerconfig

import (
	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RmqCreds               configs.RmqCreds `yaml:"rmq"`
	LogLevel               string           `yaml:"logLevel"`
	Dsn                    string           `yaml:"postgresDsn"`
	RemoveScannerFrequency int              `yaml:"removeScannerFrequency"`
	NotifyScannerFrequency int              `yaml:"notifyScannerFrequency"`
	StoragePeriod          int              `yaml:"storagePeriod"`
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
