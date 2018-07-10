package revert

import (
	"reflect"
	"testing"

	"github.com/whyming/algorithm/base/stack"
)

func TestRevert(t *testing.T) {
	type args struct {
		s *stack.Int
	}
	originStack := stack.NewStackInt()
	originStack.Push(1)
	originStack.Push(2)
	originStack.Push(3)
	originStack.Push(4)

	revertStack := stack.NewStackInt()
	revertStack.Push(4)
	revertStack.Push(3)
	revertStack.Push(2)
	revertStack.Push(1)

	tests := []struct {
		name string
		args args
		want *stack.Int
	}{
		{name: "onlyOne", args: args{s: originStack}, want: revertStack},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Revert(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Revert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBottom(t *testing.T) {
	type args struct {
		s *stack.Int
	}
	s1 := stack.NewStackInt()
	s1.Push(4)
	s1.Push(3)
	s1.Push(2)
	s1.Push(1)

	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1", args: args{s: s1}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBottom(tt.args.s); got != tt.want {
				t.Errorf("getBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}
