package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server Server `yaml:"server"`
	GC     GC     `yaml:"gc"`
}

type Server struct {
	Port int `yaml:"port"`
}

type GC struct {
	Percent int `yaml:"percent"`
}

func MustLoad(path string) *Config {

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}