package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go routine1(ch)
	fmt.Println(<-ch)
}

func routine1(result chan string) {
	// Über tasks wird eine Verbindung zwischen
	// routine1 und routine2 hergestellt
	tasks := make(chan chan string)
	go routine2(tasks)
	// jetzt können wir einen Channel erstellen, der
	// uns das Ergebnis zurück melden kann
	task := make(chan string)
	tasks <- task
	result <- <-task
}

func routine2(tasks chan chan string) {
	for resCh := range tasks {
		// ...
		time.Sleep(time.Second)
		resCh <- "Ergebnis aus routine2"
	}
}
