package main

import (
	"fmt"
)

func main() {
	// ungerichteter Channel
	c := make(chan int)

	// in c kann jedoch geschrieben werden
	go foo(c)

	// warten auf den Wert aus c
	fmt.Println(<-c)
}

func foo(ch chan<- int) {
	// in ch kann nur geschrieben werden
	ch <- 42
}
