package main

import (
	"fmt"
)

// Ausdruck nach switch

// https://play.golang.org/p/xk4fccIau8Z

func main() {
	a := 3
	switch a {
	case 0:
		fmt.Println("a=0")
	case 2, 3, 4:
		// mehrere cases mit Komma getrennt
		fmt.Println("2, 3 oder 4")
	default:
		fmt.Println("alles andere")
	}
}
