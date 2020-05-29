package main

import (
	"fmt"
)

func main() {
	c := make(chan string, 1)
	for {
		select {
		case c <- "ping":
		case c <- "pong":
		case s := <-c:
			fmt.Println(s)
		}
	}
}
