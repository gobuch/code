package main

import (
	"fmt"
)

func main() {
	s := []string{"null", "eins", "zwei"}
	for i, v := range s {
		fmt.Printf("%02d: %s\n", i, v)
	}
}
