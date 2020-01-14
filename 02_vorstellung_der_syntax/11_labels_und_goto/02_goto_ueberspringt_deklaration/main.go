package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hallo")
	goto MeinLabel
	var i int
MeinLabel:
	i++
	fmt.Println(i)
}
