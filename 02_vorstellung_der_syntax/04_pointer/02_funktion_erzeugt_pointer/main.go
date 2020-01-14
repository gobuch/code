package main

import (
	"fmt"
)

func foo() *int {
	bar := 123
	return &bar
}
func main() {
	a := foo()
	fmt.Println(a, *a)
}
