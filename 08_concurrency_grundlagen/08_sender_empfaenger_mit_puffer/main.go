package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan string, 5)
	wg.Add(2)
	go routine1(ch, wg)
	go routine2(ch, wg)
	wg.Wait()
}

// routine1 ist der Sender
func routine1(ch chan string, wg *sync.WaitGroup) {
	ch <- "Hallo routine2"
	ch <- "wichtige Nachricht"
	wg.Done()
}

// routine2 ist der Empfänger
func routine2(ch chan string, wg *sync.WaitGroup) {
	// warte auf die Nachricht
	nachricht := <-ch
	fmt.Println("routine2: ", nachricht)
	wg.Done()
}
