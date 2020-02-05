package main

import "testing"

func TestValidateURL(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"https://golang.org", true},
		{"http://golang.org", true},
		{"http://golang.org/abc", true},
		{"http://golang.org/abc?id=123", true},
		{"http//golang.org", false},
		{"golang.org", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := validateURL(tt.s); got != tt.want {
				t.Errorf("validateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
