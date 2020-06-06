package sum

import (
	"bytes"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		input   string
		want    string
		wantErr bool
	}{
		{
			"[1,2,3]",
			"6",
			false,
		},
		{
			"[]",
			"0",
			false,
		},
		{
			"[a,b]",
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			b := &bytes.Buffer{}
			err := sum(r, b)
			if (err != nil) != tt.wantErr {
				t.Errorf("sum() error = %v want: %v\n", err, tt.wantErr)
				return
			}
			if b.String() != tt.want {
				t.Errorf("sum()=%s want: %s\n", b.String(), tt.want)
			}
		})
	}
}
