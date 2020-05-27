package main

import (
	"fmt"
)

// https://play.golang.org/p/Ns96QhdWVlB

func main() {
	x := 3
	// Einfache if-Abfrage
	if x > 5 {
		doThis(x)
	} else if x > 0 {
		doThat(x)
	} else {
		doSomething(x)
	}

	// if mit Deklaration
	if a := getNumber(); a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println(a)
	}
	// ab hier existiert a nicht mehr
}

func doThis(i int) {
	fmt.Println("doThis:", i)
}

func doThat(i int) {
	fmt.Println("doThis:", i)
}

func doSomething(i int) {
	fmt.Println("doSomething:", i)
}

func getNumber() int {
	return 27
}
