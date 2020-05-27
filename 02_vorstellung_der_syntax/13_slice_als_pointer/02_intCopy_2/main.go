package main

import (
	"fmt"
)

func intCopy(ziel []int, quelle []int) int {
	n := 0
	for i := range ziel {
		if i >= len(quelle) {
			break
		}
		ziel[i] = quelle[i]
		n++
	}
	return n
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int, 3)
	n := intCopy(b, a)
	fmt.Printf("%d Elemente kopiert: b=%#v", n, b)
}
