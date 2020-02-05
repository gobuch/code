package main

import (
	"flag"
	"fmt"
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
	// Ausgabe damit der Code kompiliert
	fmt.Println(url)
}
