package main

type intStack struct {
	stackSlice []int
}

func (s *intStack) Push(in int) {
	s.stackSlice = append(s.stackSlice, in)
}

func (s *intStack) Pop() int {
	if len(s.stackSlice)==0{
		var empty int
		return empty
	}
	out := s.stackSlice[len(s.stackSlice)-1]
	s.stackSlice = s.stackSlice[:len(s.stackSlice)-1]
	return out
}
