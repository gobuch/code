package main

import (
	"reflect"
	"testing"
)

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []cmdArgs
	}{
		{
			"ein Befehl",
			[]string{"abc", "def", "hij"},
			[]cmdArgs{
				cmdArgs{"abc", []string{"def", "hij"}},
			},
		},
		{
			"zwei Befehle",
			[]string{"abc", "def", "::", "abcd", "defg"},
			[]cmdArgs{
				cmdArgs{"abc", []string{"def"}},
				cmdArgs{"abcd", []string{"defg"}},
			},
		},
		{
			"kein Befehl",
			[]string{},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseArgs(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
