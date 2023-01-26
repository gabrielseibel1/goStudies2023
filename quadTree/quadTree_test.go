package quadtree

import (
	"reflect"
	"testing"
)

func TestQuadTree(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "[]=>[]",
			args: args{
				grid: [][]int{},
			},
			want: nil,
		},
		{
			name: "[[0,1][1,0]]=>[[0,0], [0,1], [1,1], [1,1], [0,1]]",
			args: args{
				grid: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: &Node{
				Val:    false,
				IsLeaf: false,
				TopLeft: &Node{
					Val:    false,
					IsLeaf: true,
				},
				TopRight: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomLeft: &Node{
					Val:    true,
					IsLeaf: true,
				},
				BottomRight: &Node{
					Val:    false,
					IsLeaf: true,
				},
			},
		},
		{
			name: "[[1,1][1,1]]=>[[1,1]]",
			args: args{
				grid: [][]int{
					{1, 1},
					{1, 1},
				},
			},
			want: &Node{
				Val:    true,
				IsLeaf: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuadTree(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuadTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
