package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var r io.Reader
	r = bytes.NewBufferString("Hallo Interface")
	f := r.(*os.File)
	defer f.Close()
}
