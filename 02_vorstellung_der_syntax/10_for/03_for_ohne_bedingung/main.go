package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		if i > 4 {
			break
		}
		fmt.Println(i)
		i++
	}
}
