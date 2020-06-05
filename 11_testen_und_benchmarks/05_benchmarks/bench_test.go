package benchmarks

import (
	"fmt"
	"testing"
)

var s string

var (
	x = "abcd"
	y = "abcd"
)

func BenchmarkSprintf(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = fmt.Sprintf("%s%s", x, y)
	}
	s = str
}

func BenchmarkPlus(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = x + y
	}
	s = str
}
