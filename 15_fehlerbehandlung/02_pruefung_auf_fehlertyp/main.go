package main

import (
	"errors"
	"fmt"
)

type barError string

func (e barError) Error() string {
	return string(e)
}

func main() {
	err := foo()

	fmt.Println(err)
	var e barError
	fmt.Println(errors.As(err, &e))
	fmt.Println(e)
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
	return barError("barFehler")
}

// Output:
// bar(): barbar(): barFehler
// true barFehler
