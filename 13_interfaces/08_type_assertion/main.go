package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var r io.Reader
	r = bytes.NewBufferString("Hallo Interface")
	b := r.(*bytes.Buffer)
	fmt.Println(b.Bytes())
}
