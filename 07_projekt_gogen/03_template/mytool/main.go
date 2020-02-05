package main

import "fmt"

type MeinTyp struct {
	Foo string
	Bar string
}

func main() {
	mts := MeinTypStack{}
	mts.Push(MeinTyp{"foo1", "bar1"})
	mts.Push(MeinTyp{"foo2", "bar2"})
	p := mts.Pop()
	fmt.Printf("%#v\n", p)
	is := intStack{}
	is.Push(1)
	is.Push(2)
	is.Push(3)
	fmt.Println(is.Pop())
	fmt.Println(is.Pop())

}
