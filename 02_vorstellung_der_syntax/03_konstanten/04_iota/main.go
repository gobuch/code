package main

import "fmt"

// https://play.golang.org/p/ANpWTiSlR9p

const (
	Montag     = iota // 0
	Dienstag   = iota // 1
	Mittwoch   = iota // 2
	Donnerstag        // 3
	Freitag           // 4
	Samstag           // 5
	Sonntag           // 6
)

// Beginnt hier wieder bei 0
const (
	_ = iota // 0 auslassen
	eins
	_ // 2 auslassen
	drei
)

type ByteSize float64

const (
	_           = iota // ignoriere 0
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// Beispiele für die Tage
	fmt.Println(Montag, Sonntag)

	fmt.Println(eins, drei)

	// Anzahl Bytes für GB
	fmt.Println(GB)
}
