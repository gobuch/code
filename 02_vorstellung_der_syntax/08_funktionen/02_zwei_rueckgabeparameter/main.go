package main

import (
	"fmt"
)

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := swap(12, 2)
	fmt.Println(a)
	fmt.Println(b)
}
