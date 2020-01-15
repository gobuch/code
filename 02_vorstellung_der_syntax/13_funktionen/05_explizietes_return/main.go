package main

import (
	"fmt"
)

type item struct {
	name  string
	value int
}

func findItem(items []item, name string) (result item, ok bool) {
	for _, i := range items {
		if i.name == name {
			result = i
			return result, true
		}
	}
	return result, false
}

func main() {
	liste := []item{
		item{"Alice", 10},
		item{"Bob", 8},
		item{"Chris", 14},
	}
	result, ok := findItem(liste, "Bob")
	if ok {
		fmt.Println(result)
	}

}
