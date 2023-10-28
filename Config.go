package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	AppConfig struct {
		DatabaseUrl string `yaml:"database_url"`
	}

	ConfigFile map[string]*AppConfig
)

func LoadConfig(env string) (*AppConfig, error) {
	configFile := ConfigFile{}
	file, _ := os.Open("config.yml")
	defer file.Close()
	decoder := yaml.NewDecoder(file)

	// Always check for errors!
	if err := decoder.Decode(&configFile); err != nil {
		return nil, err
	}

	appConfig, ok := configFile[env]
	if !ok {
		return nil, fmt.Errorf("no such environment: %s", env)
	}

	return appConfig, nil
}
