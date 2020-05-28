package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(c chan int) {
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c)
	}(ch)
	for i := 0; i < 4; i++ {
		n, ok := <-ch
		fmt.Println(n, ok)
	}
}
