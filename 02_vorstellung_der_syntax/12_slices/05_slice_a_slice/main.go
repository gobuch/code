package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a[1])   // 1
	fmt.Println(a[0:3]) // [0 1 2]
	fmt.Println(a[:3])  // [0 1 2]
	fmt.Println(a[3:5]) // [3 4]
	fmt.Println(a[3:])  // [3 4 5]
	fmt.Println(a[:])   // [0 1 2 3 4 5]
}
