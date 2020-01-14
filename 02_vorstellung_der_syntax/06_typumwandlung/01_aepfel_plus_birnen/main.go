package main

import (
	"fmt"
)

/*
Code kompiliert nicht, da wir Äpfel mit Birnen
vergleichen.
*/

func main() {
	type äpfel int
	type birnen int
	a := äpfel(10)
	b := birnen(5)
	fmt.Println(a + b)
}
