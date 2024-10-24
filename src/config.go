package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Embed struct {
	Title       string `yaml:"Title"`
	Description string `yaml:"Description"`
	URL         string `yaml:"Url,omitempty"`
	Color       int    `yaml:"Color"`
}

type Config struct {
	TOPGG_KEY   string `yaml:"TOPGG_KEY"`
	WEBHOOK_URL string `yaml:"WEBHOOK_URL"`
	PORT        int    `yaml:"PORT"`
	ADDRESS     string `yaml:"ADDRESS"`
	EMBED       Embed  `yaml:"EMBED"`
	ENDPOINT    string `yaml:"ENDPOINT"`
}

func Init(filename string) (Config, error) {
	var config Config

	// Read the YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
