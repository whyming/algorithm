package stack

import "errors"

// Int is a stack of int
type Int []int

// NewStackInt return a StatckInt instance
func NewStackInt() *Int {
	return &Int{}
}

// IsEmpty ...
func (s Int) IsEmpty() bool {
	return len(s) <= 0
}

// Push an item to stack
func (s *Int) Push(d int) {
	*s = append(*s, d)
}

// Pop an the peek data
func (s *Int) Pop() (result int, err error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	result = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

// Peek return the peek data
func (s Int) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}
	return s[len(s)-1], nil
}
