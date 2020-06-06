package main

import (
	"fmt"
)

func main() {
	saveFoo(1)
	saveFoo(10)
	saveFoo(2)
}

func foo(i int) {
	a := []int{0, 1, 2}
	fmt.Printf("foo(%d): %d\n", i, a[i])
}

func saveFoo(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover() ausgeführt")
			fmt.Printf("foo(%d): err: %s\n", i, err)
		}
	}()
	foo(i)
}
