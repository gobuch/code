package main

import (
	"fmt"
)

func printSlice(name string, slice ...[]int) {
	fmt.Println(name, ":")
	for _, s := range slice {
		fmt.Printf("%p - len: %d cap: %d - %#v\n", s, len(s), cap(s), s)
	}
}

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[1:5]
	printSlice("b := a[1:5]", a, b)
	b = append(b, 11)
	printSlice("nach append", a, b)
	c := a[1:5:5]
	printSlice("c := a[1:5:5]", c)
	c = append(c, 14)
	printSlice("nach c append", a, c)
}
