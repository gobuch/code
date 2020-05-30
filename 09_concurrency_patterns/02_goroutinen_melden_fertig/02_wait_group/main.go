package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		doSomething()
		wg.Done()
	}()
	fmt.Println("warte")
	wg.Wait()
	fmt.Println("fertig")
}

func doSomething() {
	// irgendwas langsames
	time.Sleep(time.Second * 2)
}
