package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func initConfig(configPath string) error {
	log.Printf("Reading the configuration\n%s\n", configPath)
	yamlData, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	log.Printf("Got yaml data\n%s\n", yamlData)
	onceCfg.Do(func() {
		config, err = newConfig(&yamlData)
	})
	if err != nil {
		return err
	}
	return nil
}

func newConfig(yamlData *[]byte) (*Config, error) {
	var c *Config
	err := yaml.Unmarshal(*yamlData, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
