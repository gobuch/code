package main

import (
	"fmt"
)

// https://play.golang.org/p/DdRHXQVjEkj

type adresse struct {
	strasse string
	stadt   string
}

func main() {
	// Deklaration mit Feldnamen
	a := adresse{
		strasse: "Musterstr.",
		stadt:   "Musterstadt",
	}

	// Deklaration mit allen Feldern
	b := adresse{"Bahnhofstr.", "Berlin"}

	fmt.Println(a)
	fmt.Println(b)
}
