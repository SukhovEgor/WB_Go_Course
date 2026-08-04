package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server  ServerConfig  `yaml:"server"`
	Cluster ClusterConfig `yaml:"cluster"`
	Workers int           `yaml:"workers"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type ClusterConfig struct {
	Nodes []string `yaml:"nodes"`
}

func MustLoad(path string) *Config {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	cfg := &Config{}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		panic(err)
	}

	return cfg
}