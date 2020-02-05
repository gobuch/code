package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	var w io.Writer
	w = os.Stdout
	if *flagOutput != "" {
		err := os.MkdirAll(filepath.Dir(*flagOutput), 0755)
		if err != nil {
			fmt.Printf("Fehler beim Anlegen des Ordners: %v", err)
			os.Exit(1)
		}
		f, err := os.OpenFile(*flagOutput, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			fmt.Printf("Fehler beim Anlegen von %s\n%#v", *flagOutput, err)
		}
		defer f.Close()
		w = f

	}
	io.Copy(w, resp.Body)
}
