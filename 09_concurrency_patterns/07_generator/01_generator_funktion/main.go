package main

import "fmt"

func gen() chan int {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			ch <- i
			i++
			// hier im Beispiel brechen wir
			// bei 5 ab. Abbruchbedingung ist
			// so nicht im Buch
			if i > 5 {
				// ab hier werden keine Daten
				// gesendet, also schließen wir
				// den Channel
				close(ch)
				break
			}
		}
	}()
	return ch
}

func main() {
	ch := gen()
	for nr := range ch {
		fmt.Println(nr)
	}
}
