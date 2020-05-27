package main

import (
	"fmt"
)

// https://play.golang.org/p/PvUEphpkUPJ

func main() {
	a := 0
	switch {
	case a == 0:
		fmt.Println("a=0")
	default:
		fmt.Println("alles andere")
	}

}
