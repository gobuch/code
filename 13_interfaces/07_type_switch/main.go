package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var r io.Reader
	myFunc(r)
	r = bytes.NewBufferString("Hallo Interface")
	myFunc(r)
}

func myFunc(r io.Reader) {
	switch v := r.(type) {
	case nil:
		fmt.Println("nil Pointer")
	case *bytes.Buffer:
		fmt.Println("bytes.Buffer", v.Bytes())
	default:
		fmt.Println("Typ unbekannt")
	}
}
