package main

import (
	"fmt"
)

type filterFunc func(string) bool

func meinFilter(s []string, filter filterFunc) []string {
	var out []string
	for _, str := range s {
		if filter(str) {
			out = append(out, str)
		}
	}
	return out
}

func laengenFilter(laenge int) filterFunc {
	return func(s string) bool {
		if len(s) < laenge {
			return false
		}
		return true
	}
}

func main() {
	f3 := laengenFilter(3)
	f2 := laengenFilter(2)
	s := []string{"a", "abcd", "abc", "ab"}
	fmt.Println(meinFilter(s, f2))
	fmt.Println(meinFilter(s, f3))
	// Output:
	// [abcd abc ab]
	// [abcd abc]
}
