package main

import "fmt"

func gen(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case ch <- i:
				i++
			case <-done:
				close(ch)
				return
			}
		}
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := gen(done)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(done)
}
