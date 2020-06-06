// Package structfield is used to handle different field

// in structs

package structfield

import (
	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name       string
		dst        interface{}
		src        interface{}
		wantErr    bool
		wantString string
	}{
		{
			"Testfall nur B",
			&struct {
				A string
				B string
			}{"-", "-"},
			struct {
				B string
				C string
			}{"B", "C"},
			false,
			"&{- B}",
		},
		{
			"Testfall kein Pointer",
			struct {
				A string
				B string
			}{"-", "-"},
			struct {
				B string
				C string
			}{"B", "C"},
			true,
			"{- -}",
		},
		{
			"Testfall anderer Feldtyp",
			&struct {
				A string
				B int
			}{"-", 5},
			struct {
				B string
				C string
			}{"B", "C"},
			false,
			"&{- 5}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Copy(tt.dst, tt.src); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if fmt.Sprintf("%v", tt.dst) != tt.wantString {
				t.Errorf("Got: \n%v\nWant:\n%s", tt.dst, tt.wantString)
			}
		})
	}
}
