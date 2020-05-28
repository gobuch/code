package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5; i++ {
		go zuhörer(ch)
		time.Sleep(time.Second)
	}
	ch <- "hab nur eine Nachricht"
	time.Sleep(time.Second)
	fmt.Println("Ende")
}

func zuhörer(ch chan string) {
	fmt.Println("warte ...")
	nachricht := <-ch
	fmt.Println("Nachricht erhalten: ", nachricht)
}
