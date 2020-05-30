package main

import "fmt"

func main() {
	done := make(chan struct{})
	select {
	case <-done:
		panic("Channel ist geschlossen")
	default:
	}
	fmt.Println("Hallo zusammen!")
}
