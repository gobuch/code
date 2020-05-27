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

func main() {
	f := func(s string) bool {
		if len(s) < 3 {
			return false
		}
		return true
	}
	s := []string{"a", "abcd", "abc", "ab"}
	fmt.Println(meinFilter(s, f))
	// Output: [abcd abc]
}
