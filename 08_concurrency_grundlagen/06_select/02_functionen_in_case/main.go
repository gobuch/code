package main

import "fmt"

func main() {
	c := make(chan int, 1)
	counter := 0
	count := func(i int) int {
		counter++
		return i
	}
	select {
	case c <- count(1):
	case c <- count(2):
	case c <- count(3):
	case c <- count(4):
	case c <- count(5):
	}
	fmt.Println("c:", <-c)
	fmt.Println("counter:", counter)
}
