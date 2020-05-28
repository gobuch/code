package main

import (
	"fmt"
	"sync"
)

// https://play.golang.org/p/NguvPUBkeO-

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(name string) {
		fmt.Printf("Hallo %s", name)
		wg.Done()
	}("Till")
	wg.Wait()
}
