package stack

import (
	"testing"
)

func TestIntStack(t *testing.T) {
	s := intStack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	for i := 4; i > 0; i-- {
		got := s.Pop()
		want := i
		if got != want {
			t.Errorf("pop()= %v, want: %v", got, want)
		}
	}
}
