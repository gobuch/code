package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}, 1)
	go func() {
		doSomething()
		done <- struct{}{}
	}()

	fmt.Println("warte")
	<-done
	fmt.Println("fertig")
}

func doSomething() {
	// irgendwas langsames
	time.Sleep(time.Second * 2)
}
