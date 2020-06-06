package main

import (
	"fmt"
	"io"
)

func main() {
	var w io.Writer
	fmt.Println(w) // <nil>
}
