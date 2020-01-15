package main

import (
	"fmt"
)

func greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}

func main() {
	greeting := greet("Alice")
	fmt.Println(greeting) // Output: Hallo Alice!
}