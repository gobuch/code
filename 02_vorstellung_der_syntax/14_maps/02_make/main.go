package main

import "fmt"

func main() {
	var m map[int]string
	m = make(map[int]string)
	m[1] = "Eintrag 1"
	fmt.Println(m)
}
