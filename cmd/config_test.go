package main

import (
	"fmt"
	"testing"
)

func TestInitConfig(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  string
	}{
		{"Right path", "../configs/config.yaml", "<nil>"},
		{"Wrong path", "config.yaml", "open config.yaml: no such file or directory"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out = fmt.Sprintf("%v", initConfig(tt.input))
			if out != tt.want {
				t.Errorf("got \"%v\"\nwant \"%v\"", out, tt.want)
			}
		})
	}
}
