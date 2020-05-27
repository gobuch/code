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
	e := make([]int, 2, 4)
	n := copy(e, a[1:5])
	fmt.Println("kopierte Elemente:", n)
	printSlice("copy(e,a[1:5])", a, e)
}
