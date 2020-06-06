package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MeinTyp string
	type MeinStruct struct {
		A string
		b string
	}

	var (
		myInt    int     = 12
		myString string  = "Hallo"
		myBool   bool    = true
		myType   MeinTyp = "MeinTyp"
		myStruct         = MeinStruct{"ABC", "bcd"}
	)
	refl(myInt)
	refl(myString)
	refl(myBool)
	refl(myType)
	refl(&myType)
	refl(myStruct)
}

func refl(i interface{}) {
	fmt.Println("Typ:", reflect.TypeOf(i))
	val := reflect.ValueOf(i)
	fmt.Println("Wert:", val, val.Kind())
}
