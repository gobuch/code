package main

import (
	"fmt"
)

// https://play.golang.org/p/z2EXuIaszsi

func main() {
	type äpfel int
	type birnen int
	a := äpfel(10)
	b := birnen(5)
	anzahl := int(a) + int(b)
	fmt.Printf("Anzahl Früchte: %d", anzahl)
}
