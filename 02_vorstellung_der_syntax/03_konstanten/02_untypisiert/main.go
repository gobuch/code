package main

import "fmt"

/*
Dieses Beispiel kompiliert nicht, da der Typ der
Variable ungleich int ist.

https://play.golang.org/p/OlA0DxtGlK1
*/

func add2(a int) int {
	return 2 + a
}

func main() {
	const c = 2.0
	// c ist untypisiert und kann hier verwendet werden
	fmt.Println(add2(c))
	// Output:
	// 4
	var v = 2.0
	fmt.Println(add2(v))
	// Output:
	// Fehler beim Kompilieren:
	// cannot use v (type float64) as type int in argument to add2
}
