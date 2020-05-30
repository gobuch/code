package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	myFunc()
	fmt.Println("myFunc() ist fertig")
	// warten, damit alle Print-Aufrufe
	// ausgegeben werden können
	time.Sleep(time.Second)
}

func myFunc() {
	ctx, cancel := context.WithCancel(
		context.Background())
	// cancel() immer ausführen
	defer cancel()
	ch := gen(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Beenden der Goroutine")
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
