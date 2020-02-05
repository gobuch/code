package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	flagOutput = flag.String("o", "", "output file")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("bitte nur eine url angeben")
		os.Exit(1)
	}
	url := args[0]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Fehler beim Lesen von %s\n%#v", url, err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
