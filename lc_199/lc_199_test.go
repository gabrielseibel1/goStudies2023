package lc199

import (
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0,2,5,8,11,13,14",
			args: args{
				root: &TreeNode{
					Val: 0, // can see
					Left: &TreeNode{
						Val:  1,
						Left: &TreeNode{Val: 3},
					},
					Right: &TreeNode{
						Val: 2, // can see
						Left: &TreeNode{
							Val: 4,
							Left: &TreeNode{
								Val: 6,
								Left: &TreeNode{
									Val:  9,
									Left: &TreeNode{Val: 12},
									Right: &TreeNode{
										Val:  13,                 // can see
										Left: &TreeNode{Val: 14}, // can see
									},
								},
							},
							Right: &TreeNode{
								Val:   7,
								Left:  &TreeNode{Val: 10},
								Right: &TreeNode{Val: 11}, // can see
							},
						},
						Right: &TreeNode{
							Val:   5,                 // can see
							Right: &TreeNode{Val: 8}, // can see
						},
					},
				},
			},
			want: []int{0, 2, 5, 8, 11, 13, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
