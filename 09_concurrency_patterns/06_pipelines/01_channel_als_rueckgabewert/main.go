package main

import "fmt"

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})
	ch := gen(done)
	quadratChannel := sq(ch)
	fmt.Println(<-quadratChannel)
	fmt.Println(<-quadratChannel)
	fmt.Println(<-quadratChannel)
	close(done)
}

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
