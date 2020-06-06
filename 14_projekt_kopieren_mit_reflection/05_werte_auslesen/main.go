package main

import (
	"fmt"
	"reflect"
)

func main() {
	myStruct := struct {
		Feld1 string
		Feld2 string
	}{
		"Wert Feld1",
		"Wert Feld2",
	}
	refl3(myStruct)
}

func refl3(i interface{}) {
	val := reflect.ValueOf(i)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println("Name:", field.Name)
		fmt.Println("Wert:", val.Field(i))
	}
}
