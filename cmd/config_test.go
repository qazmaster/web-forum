package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitConfig(t *testing.T) {
	var pathTests = []struct {
		name  string
		input string
		want  string
	}{
		{"Right path", "../configs/config.yaml", "<nil>"},
		{"Wrong path", "config.yaml", "open config.yaml: no such file or directory"},
	}
	for _, tt := range pathTests {
		t.Run(tt.name, func(t *testing.T) {
			var out = fmt.Sprintf("%v", initConfig(tt.input))
			if out != tt.want {
				t.Errorf("got \"%v\"\nwant \"%v\"", out, tt.want)
			}
		})
	}
}

func TestNewConfig(t *testing.T) {
	var (
		defaultCfgIn = `server:
  host: localhost
  port: 8080
database:
  dsn: forum.db`
		defaultCfgOut = &Config{
			Server: struct {
				Host string `yaml:"host"`
				Port int    `yaml:"port"`
			}{
				Host: "localhost",
				Port: 8080,
			},
			Database: struct {
				DSN string `yaml:"dsn"`
			}{
				DSN: "forum.db",
			},
		}
	)

	var cgfTests = []struct {
		name  string
		input []byte
		want  *Config
	}{
		{"Default config", []byte(defaultCfgIn), defaultCfgOut},
	}
	for _, tt := range cgfTests {
		t.Run(tt.name, func(t *testing.T) {
			out, _ := newConfig(&tt.input)
			if !Equal(out, tt.want) {
				t.Errorf("\ngot %v\nwant %v", out, tt.want)
			}
		})
	}
}

func Equal(a, b any) bool {
	return reflect.ValueOf(a) == reflect.ValueOf(b)
}
