package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Node struct {
	Name     string `yaml:"name"`
	Command  string `yaml:"command,omitempty"`  // Empty if folder
	Children []Node `yaml:"children,omitempty"` // Empty if command
}

type Config struct {
	Commands []Node `yaml:"commands"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
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
