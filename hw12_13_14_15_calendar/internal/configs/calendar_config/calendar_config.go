package calendarconfig

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host            string `yaml:"host"`
	Port            string `yaml:"httpPort"`
	EventServerPort string `yaml:"grpcPort"`
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
	k8s := os.Getenv("k8s")
	if k8s != "" {
		config.Dsn = "host=calendar-statefulset-0.calendar-service.default.svc.cluster.local port=5432 " +
			"user=postgres password=postgres dbname=calendar sslmode=disable"
	}
	return config, nil
}
