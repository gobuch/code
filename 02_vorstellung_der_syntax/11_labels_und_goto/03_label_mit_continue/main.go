package main

import (
	"fmt"
)

func main() {
ZeilenLabel:
	for zeile := 0; zeile <= 4; zeile++ {
		for spalte := 0; spalte <= 3; spalte++ {
			fmt.Printf("(%d,%d) ", zeile, spalte)
			if zeile+spalte == 4 {
				fmt.Printf("\n")
				continue ZeilenLabel
			}
		}
		fmt.Printf("\n")
	}
}
