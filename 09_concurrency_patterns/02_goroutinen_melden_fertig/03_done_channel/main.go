package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// ...
	go func() {
		doSomething()
		close(done)
	}()
	// Programmfluss blockiert hier, bis
	// done geschlossen wird.

	fmt.Println("warte")
	<-done
	fmt.Println("fertig")
}

func doSomething() {
	// irgendwas langsames
	time.Sleep(time.Second * 2)
}
