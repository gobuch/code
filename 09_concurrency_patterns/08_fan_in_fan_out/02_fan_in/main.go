package main

import (
	"fmt"
	"sync"
)

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Goroutine schließt den Output
	// wenn alle Input-Channels geschlossen
	// wurden.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	// Drei Generatoren erzeugen Werte über mehrere Channels
	c1 := gen()
	c2 := gen()
	c3 := gen()
	// Fan-In der Channels
	m := merge(c1, c2, c3)
	for i := range m {
		fmt.Println(i)
	}
}

func gen() chan int {
	ch := make(chan int)
	go func(c chan int) {
		for i := 1; i <= 3; i++ {
			c <- i
		}
		close(c)
	}(ch)
	return ch
}
