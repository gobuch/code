package main

import "fmt"

func main() {
	done := make(chan struct{})
	// deadlock, da die Abfrage blockiert
	_, ok := <-done
	if !ok {
		fmt.Println("Channel ist geschlossen")
	}
}
