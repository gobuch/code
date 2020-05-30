package main

import (
	"fmt"
)

type resultType int

func main() {
	result := make(chan resultType, 1)
	// ...
	go func() {
		r := doSomething()
		result <- r
	}()
	// ...
	// Verwende das Ergebnis
	ergebnis := <-result
	fmt.Println(ergebnis)
}

func doSomething() resultType {
	return 5
}
