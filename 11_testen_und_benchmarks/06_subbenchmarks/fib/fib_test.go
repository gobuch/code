package fib

import (
	"fmt"
	"testing"
)

var result int

func benchmarkFib(x int, b *testing.B) {
	var f int
	for i := 0; i < b.N; i++ {
		f = fib(x)
	}
	result = f
}

func BenchmarkFibonacci(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		benchmarkFib(1, b)
	})
	b.Run("10", func(b *testing.B) {
		benchmarkFib(10, b)
	})
	b.Run("20", func(b *testing.B) {
		benchmarkFib(20, b)
	})

}

func BenchmarkFibonacciLoop(b *testing.B) {
	inValues := []int{1, 10, 20}
	for _, v := range inValues {
		b.Run(fmt.Sprintf("fib(%d)", v), func(b *testing.B) {
			var f int
			for i := 0; i < b.N; i++ {
				f = fib(v)
			}
			result = f
		})
	}

}
