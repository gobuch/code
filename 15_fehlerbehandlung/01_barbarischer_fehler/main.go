package main

import (
	"errors"
	"fmt"
)

var ErrBarBar = errors.New("barbarischer Fehler")

func main() {
	err := foo()

	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrBarBar))
}

func foo() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("bar(): %w", err)
	}
	return nil
}

func bar() error {
	err := barbar()
	if err != nil {
		return fmt.Errorf("barbar(): %w", err)
	}
	return nil
}

func barbar() error {
	return ErrBarBar
}
