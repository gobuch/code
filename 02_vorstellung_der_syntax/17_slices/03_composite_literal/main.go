package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("%p - len: %d cap: %d %#v\n",
		s, len(s), cap(s), s)
}

func main() {
	// Composite Literal
	b := []int{}
	printSlice(b)
	c := []int{1, 2, 3}
	printSlice(c)

	d := make([]int, 2, 8)
	printSlice(d)
}
