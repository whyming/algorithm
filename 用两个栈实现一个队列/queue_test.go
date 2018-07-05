package queue

import (
	"testing"
)

func TestQueue_Pop(t *testing.T) {

	tests := []struct {
		name       string
		init       func(*Queue)
		wantResult int
		wantErr    bool
	}{
		{
			name: "pushData", init: func(q *Queue) {

			}, wantResult: 0, wantErr: true,
		}, {
			name: "pushData", init: func(q *Queue) {
				for _, d := range []int{4, 5, 6} {
					q.Push(d)
				}
			}, wantResult: 4, wantErr: false,
		}, {
			name: "complex", init: func(q *Queue) {
				for _, d := range []int{6, 7, 8} {
					q.Push(d)
				}
				q.Pop()
				q.Pop()
				q.Push(1)
				q.Push(2)
			}, wantResult: 8, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := NewQueue()
			tt.init(q)
			gotResult, err := q.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult != tt.wantResult {
				t.Errorf("Queue.Pop() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
