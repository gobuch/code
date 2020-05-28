package main

import (
	"fmt"
	"sync"
)

// https://play.golang.org/p/lYM-iaxKjLp

func greeter(wg *sync.WaitGroup, name string) {
	fmt.Printf("Hallo %s", name)
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go greeter(wg, "Till")
	wg.Wait()
}
