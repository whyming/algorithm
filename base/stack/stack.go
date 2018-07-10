package stack

import "errors"

// StackInt is a stack of int
type StackInt []int

// NewStackInt return a StatckInt instance
func NewStackInt() *StackInt {
	return &StackInt{}
}

// IsEmpty ...
func (s StackInt) IsEmpty() bool {
	return len(s) <= 0
}

// Push an item to stack
func (s *StackInt) Push(d int) {
	*s = append(*s, d)
}

// Pop an the peek data
func (s *StackInt) Pop() (result int, err error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	result = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

// Peek return the peek data
func (s StackInt) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s[len(s)-1], nil
}
