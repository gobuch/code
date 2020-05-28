package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func raumschiff(in chan int, out chan int) {
	a := <-in
	b := <-in
	out <- add(a, b)
}

func main() {
	in := make(chan int)
	out := make(chan int)
	go raumschiff(in, out)
	in <- 1
	in <- 2
	fmt.Println(<-out)
}
