package main

import (
	"reflect"
	"testing"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		s       string
		want    picSize
		wantErr bool
	}{
		{"500x500", picSize{500, 500}, false},
		{"500-500", picSize{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, err := parseSize(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSize(%s) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSize(%s)=%v want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestUseFile(t *testing.T) {
	tests := []struct {
		filename string
		want     bool
	}{
		{"a/path/file.JPG", true},
		{"file.jpeg", true},
		{"pib.ABC", false},
		{"justfile", false},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			if got := useFile(tt.filename); got != tt.want {
				t.Errorf("useFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
