package calc

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	r := Add(2, 3)
	fmt.Println(r)
	//Output: 4
}

func TestAdd(t *testing.T) {
	tests := []struct {
		testfall string
		a        int
		b        int
		erwartet int
	}{
		{
			"Testfall Add(1, 2)",
			1,
			2,
			3,
		},
		{
			"Testfall Add(10, 20)",
			10,
			20,
			30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testfall, func(t *testing.T) {
			s := Add(tt.a, tt.b)
			if s != tt.erwartet {
				t.Errorf("Add(%d, %d) = %d", tt.a, tt.b, s)
			}
		})
	}
}
