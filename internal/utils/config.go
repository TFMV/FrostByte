package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Polaris struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
	} `yaml:"polaris"`
}

func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
