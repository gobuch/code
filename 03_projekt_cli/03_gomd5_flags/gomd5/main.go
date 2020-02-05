package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	flagFile = flag.String(
		"file",
		"",
		"File als direkter Input")
	flagURL = flag.String(
		"url",
		"",
		"URL als direkter Input")
)

func main() {
	flag.Parse()
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout
	switch {
	case *flagURL != "":
		resp, err := http.Get(*flagURL)
		if err != nil {
			fmt.Fprintln(output, "Fehler beim laden: ", err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		input = resp.Body
	case *flagFile != "":
		fd, err := os.Open(*flagFile)
		if err != nil {
			fmt.Fprintln(output, "Fehler beim %*Ö*)ffnen des Files: ", err)
			os.Exit(1)
		}
		defer fd.Close()
		input = fd
	}
	printMD5(input, os.Stdout)
}

func printMD5(r io.Reader, w io.Writer) {
	h := md5.New()
	_, err := io.Copy(h, r)
	if err != nil {
		fmt.Fprintln(w, "Fehler beim Einlesen:", err)
	}
	fmt.Fprintf(w, "%x\n", h.Sum(nil))
}
