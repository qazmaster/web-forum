package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config *Config

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

func initConfig(configPath string) error {
	log.Printf("Reading the configuration from the %s\n", configPath)
	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	log.Printf("Got yaml data from the config file \n%s\n", yamlData)

	config = &Config{}
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return err
	}

	return nil
}
