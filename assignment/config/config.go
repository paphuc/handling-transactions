package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ReadConfig(path string) (*Config, error) {
	var c *Config
	conf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type (
	Config struct {
		GRPC   GRPC   `yaml:"grpc"`
		Server Server `yaml:"server"`
	}

	GRPC struct {
		Address string `yaml:"address"`
	}

	Server struct {
		Port int `yaml:"port"`
	}
)
