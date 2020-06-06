package main

import (
	"fmt"
	"strings"
)

type Stringer interface {
	String() string
}

func Ausgabe(s Stringer) {
	fmt.Println(s.String())
}

func GrosseAusgabe(s Stringer) {
	upper := strings.ToUpper(s.String())
	fmt.Println(upper)
}

func main() {
	type meinTyp string
	var a meinTyp = "abc"
	Ausgabe(a)
}
