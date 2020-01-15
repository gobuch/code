package main

import (
	"fmt"
)

func meineFunktion() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func main() {
	meineFunktion()
}

// Output:
// 3
// 2
// 1
