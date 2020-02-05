package main

// https://play.golang.org/p/90w-M9Cu2g2

import (
	"fmt"
)

func printSlice(s []int) {
	fmt.Printf("%p - len: %d cap: %d %#v\n",
		s, len(s), cap(s), s)
}

func main() {
	var a []int
	printSlice(a)
	a = append(a, 1)
	printSlice(a)
	a = append(a, 2)
	printSlice(a)
	a = append(a, 3)
	printSlice(a)
}
