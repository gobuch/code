package main

import (
	"fmt"
)

func main() {
	for zeile := 0; zeile <= 2; zeile++ {
		for spalte := 0; spalte <= 3; spalte++ {
			fmt.Printf("(%d,%d) ", zeile, spalte)
			if zeile+spalte == 4 {
				goto MeinLabel
			}
		}
		fmt.Printf("\n")
	}
MeinLabel:
	fmt.Println("Mein Label")
}
