package revert

import (
	"github.com/whyming/algorithm/base/stack"
)

// Revert ...
func Revert(s *stack.StackInt) *stack.StackInt {
	if s.IsEmpty() {
		return s
	}
	bottom := getBottom(s)
	Revert(s)
	s.Push(bottom)
	return s
}

func getBottom(s *stack.StackInt) int {
	peek, _ := s.Peek()
	s.Pop()
	if s.IsEmpty() {
		return peek
	}
	bottom := getBottom(s)
	s.Push(peek)
	return bottom
}
