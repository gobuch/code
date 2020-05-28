package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 1           // schreibe Wert in Puffer
	ch <- 2           // schreibe Wert in Puffer
	fmt.Println(<-ch) // aus Puffer lesen
	fmt.Println(<-ch) // aus Puffer lesen
	fmt.Println("Ende")
}
