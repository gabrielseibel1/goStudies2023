package sortedlisttobst

import (
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/list"
)

func TestSortedListToBST(t *testing.T) {
	type args struct {
		head *list.Node
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[]=>[]",
			args: args{
				head: list.New([]int{}),
			},
			want: nil,
		},
		{
			name: "[-10,-3,0,5,9]=>[0,-3,9,-10,null,5]",
			args: args{
				head: list.New([]int{-10, -3, 0, 5, 9}),
			},
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
