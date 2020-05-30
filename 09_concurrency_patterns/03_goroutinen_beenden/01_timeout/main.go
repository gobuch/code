package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan []byte, 1)
	go func() {
		result <- networkRequest()
	}()
	select {
	case data := <-result:
		fmt.Println(data)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}

func networkRequest() []byte {
	r := []byte{}
	// sehr langsamer Request
	time.Sleep(5 * time.Second)
	return r
}
