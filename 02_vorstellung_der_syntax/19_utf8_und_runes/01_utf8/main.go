package main

import (
	"fmt"
)

func main() {
	world := "世界"
	fmt.Println(len(world))
	for i, v := range world {
		fmt.Println(i, v)
	}
}
