package main

import (
	"fmt"
)

// https://play.golang.org/p/oAT_jhzsAII

type (
	meter      int
	zentimeter int
)

func MeterToZentimeter(m meter) zentimeter {
	return zentimeter(m * 100)
}

func main() {
	var a zentimeter
	var b meter = 1
	a = MeterToZentimeter(b)

	fmt.Println(a)
}
