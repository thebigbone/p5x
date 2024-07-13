package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Node struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"node"`
}

func parseConfig(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
