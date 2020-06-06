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

type user struct {
	name string
}

func (u user) String() string {
	return "Name: " + u.name
}

func main() {
	gopher := user{"Gopher"}
	Ausgabe(gopher)
	GrosseAusgabe(gopher)
}
