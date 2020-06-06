package main

type Form interface {
	Fläche() int
	Position() (int, int)
}

type Rechteck struct {
	länge  int
	breite int
	x      int
	y      int
}

func (r Rechteck) Fläche() int {
	return r.länge * r.breite
}

func (r Rechteck) Position() (int, int) {
	// ...
	return r.x, r.y
}

type Flächer interface {
	Fläche() int
}

func gleicheFläche(a, b Flächer) bool {
	return a.Fläche() == b.Fläche()
}

func main() {

}
