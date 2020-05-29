package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Diesen Code mit dem Race Detector ausführen!
// go run -race main.go

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
