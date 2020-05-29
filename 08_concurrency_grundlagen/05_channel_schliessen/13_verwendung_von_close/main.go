package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(chan int) {
		for i := 0; i < 2; i++ {
			ch <- i
		}
		close(ch)
	}(ch)
	for {
		n, ok := <-ch
		if !ok {
			break
		}
		// mach was mit n
		fmt.Println(n, ok)
	}
}
