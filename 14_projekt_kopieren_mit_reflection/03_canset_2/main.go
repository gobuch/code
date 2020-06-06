package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 100
	refl2(myInt)
	refl2(&myInt)
}

func refl2(i interface{}) {
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	fmt.Println(val.CanSet())
}
