package main

import (
	"fmt"
)

type task struct {
	a, b   int
	result chan int
}

func main() {
	// Goroutine 1
	// Channel für routine2 erzeugen
	ch := make(chan task)
	go routine2(ch)
	// Channel für das Ergebnis
	resCh := make(chan int)
	ch <- task{a: 3, b: 4, result: resCh}
	close(ch) // keine weiteren tasks mehr
	fmt.Println(<-resCh)
}

func routine2(tasks chan task) {
	for task := range tasks {
		r := task.a + task.b
		task.result <- r
	}
}
