package senderconfig

import (
	"os"

	"github.com/NoisyPunk/otus_go_hw/hw12_13_14_15_calendar/internal/configs"
	"gopkg.in/yaml.v3"
)

type Config struct {
	RmqCreds configs.RmqCreds `yaml:"rmq"`
	LogLevel string           `yaml:"logLevel"`
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
	k8s := os.Getenv("k8s")
	if k8s != "" {
		config.RmqCreds.Host = "calendar-statefulset-0.calendar-service.default.svc.cluster.local"
	}
	return config, nil
}
