package list

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "[]=nil",
			args: args{
				s: []int{},
			},
			want: nil,
		},
		{
			name: "[1,2,3]=1->2->3",
			args: args{
				s: []int{1, 2, 3},
			},
			want: &Node{
				Val: 1,
				Next: &Node{
					Val: 2,
					Next: &Node{
						Val: 3,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
