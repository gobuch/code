package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w io.Writer
	b := &bytes.Buffer{}
	w = b
	w.Write([]byte("hallo"))
	w = nil
	fmt.Println(w) // <nil>
}
