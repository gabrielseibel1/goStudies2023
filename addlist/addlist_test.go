package addlist

import (
	"reflect"
	"testing"
)

func makeList(s []int) *ListNode {
	l := new(ListNode)
	n := l
	for i, v := range s {
		n.Val = v
		if i+1 < len(s) {
			n.Next = new(ListNode)
			n = n.Next
		}
	}
	return l
}

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "345+564=807",
			args: args{
				l1: makeList([]int{2, 4, 3}),
				l2: makeList([]int{5, 6, 4}),
			},
			want: makeList([]int{7, 0, 8}),
		},
		{
			name: "0+0=0",
			args: args{
				l1: makeList([]int{0}),
				l2: makeList([]int{0}),
			},
			want: makeList([]int{0}),
		},
		{
			name: "9999999+9999=10009998",
			args: args{
				l1: makeList([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: makeList([]int{9, 9, 9, 9}),
			},
			want: makeList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name: "9999+9999999=10009998",
			args: args{
				l1: makeList([]int{9, 9, 9, 9}),
				l2: makeList([]int{9, 9, 9, 9, 9, 9, 9}),
			},
			want: makeList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
