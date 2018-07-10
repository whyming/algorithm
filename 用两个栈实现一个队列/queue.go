package queue

import (
	"github.com/whyming/algorithm/base/stack"
)

// Queue ...
type Queue struct {
	PopStack  *stack.Int
	PushStack *stack.Int
}

// NewQueue ...
func NewQueue() *Queue {
	return &Queue{
		PopStack:  stack.NewStackInt(),
		PushStack: stack.NewStackInt(),
	}
}

// Pop ...
func (q *Queue) Pop() (result int, err error) {
	if q.PopStack.IsEmpty() {
		for !q.PushStack.IsEmpty() {
			item, _ := q.PushStack.Pop()
			q.PopStack.Push(item)
		}
	}
	result, err = q.PopStack.Pop()
	return
}

// Push ...
func (q *Queue) Push(d int) {
	q.PushStack.Push(d)
}
