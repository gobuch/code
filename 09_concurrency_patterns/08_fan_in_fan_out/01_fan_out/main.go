package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	useInt := func(id int, c <-chan int) {
		for i := range c {
			fmt.Println("ID:", id, "->", i)
		}
		fmt.Println("ID:", id, "ist fertig")
		wg.Done()
	}
	in := make(chan int)
	wg.Add(3)
	go useInt(1, in)
	go useInt(2, in)
	go useInt(3, in)
	for i := 0; i < 5; i++ {
		in <- i
		time.Sleep(time.Millisecond)
	}
	close(in)
	wg.Wait()
}
