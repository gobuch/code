package main

import (
	"fmt"
)

func main() {
	var a int
	var b *int // Pointer auf int
	a = 123
	b = &a // Pointer auf eine Variable
	fmt.Println(b)
	*b = 100 // Dereferenzierung
	fmt.Println(a)

}
