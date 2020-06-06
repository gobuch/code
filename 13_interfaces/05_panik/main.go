package main

import (
	"io"
)

func main() {
	var w io.Writer
	w.Write([]byte("hallo"))
}
