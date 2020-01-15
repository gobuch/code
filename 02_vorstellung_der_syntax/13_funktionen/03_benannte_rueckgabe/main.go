package main

import (
	"fmt"
)

func foo(s string) (a string, b string) {
	b = s       // keine Deklaration notwendig
	return a, b // a = ""
}

func main() {
	a, b := foo("Hallo Till!")
	fmt.Println(a) // leer
	fmt.Println(b) // Hallo Till!
}
