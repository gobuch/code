package main

import (
	"fmt"
)

func main() {
	// Channel mit make erzeugen
	ch := make(chan string)
	// Goroutine als anonyme Funktion
	go func(ch chan string) {
		// Nachricht senden
		ch <- "eine Nachricht"
	}(ch) // ch als Input
	// auf %*Rückmeldung*) warten
	nachricht := <-ch
	fmt.Println(nachricht)
}
