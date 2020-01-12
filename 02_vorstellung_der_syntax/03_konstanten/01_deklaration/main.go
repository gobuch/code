package main

import "fmt"

/*
	Diese Datei beinhaltet alle Beispiele
	zur Deklaration von Konstanten

	https://play.golang.org/p/6MsHvHkTL1e
*/

const a = 10 // einfache Deklaration

// Deklaration als Gruppe
const (
	maxBreite = 100
	maxLaenge = 100
)

const b = 10 + 4 // %*Deklaration über einen Ausdruck*)

func main() {
	fmt.Println(a)
	fmt.Println(maxBreite, maxLaenge)
}
