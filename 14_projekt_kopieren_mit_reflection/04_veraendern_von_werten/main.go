package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 100
	refl2(myInt)
	fmt.Println(myInt)
	refl2(&myInt)
	fmt.Println(myInt)
}

func refl2(i interface{}) {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.CanSet() {
		val.SetInt(99)
	}
}
