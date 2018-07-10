package stack

import (
	"reflect"
	"testing"
)

func TestAllStackInt(t *testing.T) {
	sint := NewStackInt()
	if !sint.IsEmpty() {
		t.Error("new stackint is not empty")
	}
	sint.Push(1)
	sint.Push(2)
	if d, err := sint.Peek(); err != nil || d != 2 {
		t.Errorf("stacktint [1,2] peek is 2,get %d with err %v", d, err)
	}

	if _, err := sint.Pop(); err != nil {
		t.Errorf("stackitnt [1,2] pop result is error %v", err)
	}

	if d, err := sint.Peek(); err != nil || d != 1 {
		t.Errorf("stacktint [1] peek is 1,get %d with err %v", d, err)
	}
	sint.Pop()

	if _, err := sint.Pop(); err == nil {
		t.Error("stackitnt [] pop result is not  error")
	}

	if _, err := sint.Peek(); err == nil {
		t.Errorf("stacktint [] peek is error but no err")
	}
}
func TestNewStackInt(t *testing.T) {
	tests := []struct {
		name string
		want Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStackInt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStackInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackInt_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		s    Int
		want bool
	}{
		{name: "empty", s: []int{}, want: true},
		{name: "noEmpty", s: Int{1, 2, 3}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsEmpty(); got != tt.want {
				t.Errorf("StackInt.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStackInt_Push(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		s    *Int
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Push(tt.args.d)
		})
	}
}

func TestStackInt_Pop(t *testing.T) {
	tests := []struct {
		name    string
		s       *Int
		wantErr bool
	}{
		{name: "empty", s: &Int{}, wantErr: true},
		{name: "onlyOne", s: &Int{2}, wantErr: false},
		{name: "multi", s: &Int{2, 3, 4, 5}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.s.Pop(); (err != nil) != tt.wantErr {
				t.Errorf("StackInt.Pop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStackInt_Peek(t *testing.T) {
	tests := []struct {
		name    string
		s       Int
		want    int
		wantErr bool
	}{
		{name: "empty", s: Int{1, 2, 3}, want: 3, wantErr: false},
		{name: "onlyOne", s: Int{4}, want: 4, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("StackInt.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StackInt.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
