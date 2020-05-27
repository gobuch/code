package main

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("%p - len: %d cap: %d %#v\n",
		s, len(s), cap(s), s)
}

func main() {
	a := []int{}
	a = append(a, 0, 1, 2) // [0 1 2]
	printSlice(a)
	b := []int{3, 4, 5}
	a = append(a, b...) // [0 1 2 3 4 5]
	printSlice(a)
}
