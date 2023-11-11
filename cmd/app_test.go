package main

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	var (
		tests = []struct {
			name    string
			input   *Config
			wantApp *Application
		}{
			{
				name: "Default conf",
				input: &Config{
					Server: struct {
						Host string `yaml:"host"`
						Port int    `yaml:"port"`
					}{
						Host: "",
						Port: 8080,
					},
					Database: struct {
						DSN string `yaml:"dsn"`
					}{
						DSN: "forum.db",
					},
				},
				wantApp: &Application{},
			},
		}
	)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outApp, _ := newApp(tt.input)
			if outApp != tt.wantApp {
				t.Errorf("got \"%v\"\nwant \"%v\"", outApp, tt.wantApp)
			}
		})
	}
}
