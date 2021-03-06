package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.Errorf("Add(1, 2) = %d", s)
	}
}

func TestAdd2(t *testing.T) {
	s := Add(10, 20)
	if s != 30 {
		t.Errorf("Add(10, 20) = %d", s)
	}
}
