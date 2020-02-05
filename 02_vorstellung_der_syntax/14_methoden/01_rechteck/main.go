package main

import (
	"fmt"
)

type rechteck struct {
	laenge int
	breite int
}

// flaeche ist eine Methode für den Typ
// rechteck
func (r rechteck) flaeche() int {
	return r.laenge * r.breite
}

// setBreite ist eine Methode auf den
// Pointer
func (r *rechteck) setBreite(b int) {
	r.breite = b
}

// flaeche ist hier eine Funktion, die
// ein rechteck als Input erwartet
func flaeche(r rechteck) int {
	return r.laenge * r.breite
}

func main() {
	r := rechteck{laenge: 10, breite: 5}
	fmt.Println("Methodenaufruf: ", r.flaeche())
	fmt.Println("Funktionsaufruf: ", flaeche(r))
	r.setBreite(100)
	fmt.Println(r.flaeche())
}
