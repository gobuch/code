package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var r io.Reader
	r = bytes.NewBufferString("Hallo Interface")

	f, ok := r.(*os.File)
	if !ok {
		fmt.Printf("*os.File erwartet. %T bekommen", r)
		// ...
	}
	defer f.Close()
}
