package main

import (
	"fmt"
)

// https://play.golang.org/p/dpvxssiE0RQ

func main() {
	a := 1
	// b wird per Kurzdeklaration für den Block definiert
	switch b := a - 1; b {
	case 0:
		fmt.Println("b=0")
	default:
		fmt.Println("alles andere")
	}
	// b ist hier nicht mehr gültig

}
