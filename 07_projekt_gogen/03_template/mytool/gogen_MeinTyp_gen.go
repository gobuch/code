package main

type MeinTypStack struct {
	stackSlice []MeinTyp
}

func (s *MeinTypStack) Push(in MeinTyp) {
	s.stackSlice = append(s.stackSlice, in)
}

func (s *MeinTypStack) Pop() MeinTyp {
	if len(s.stackSlice)==0{
		var empty MeinTyp
		return empty
	}
	out := s.stackSlice[len(s.stackSlice)-1]
	s.stackSlice = s.stackSlice[:len(s.stackSlice)-1]
	return out
}
