package topological

import (
	"testing"
)

func TestIsTopologicallySorted(t *testing.T) {
	type args struct {
		graph   [][]int
		sorting []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "square-true",
			args: args{
				graph:   [][]int{{1, 2}, {3}, {3}, {}},
				sorting: []int{0, 2, 1, 3},
			},
			want: true,
		},
		{
			name: "square-false",
			args: args{
				graph:   [][]int{{1, 2}, {3}, {3}, {}},
				sorting: []int{0, 1, 3, 2},
			},
			want: false,
		},
		{
			name: "trapezoid-true",
			args: args{
				graph:   [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
				sorting: []int{0, 1, 2, 3, 4},
			},
			want: true,
		},
		{
			name: "trapezoid-false",
			args: args{
				graph:   [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
				sorting: []int{4, 2, 3, 1, 0},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTopologicallySorted(tt.args.graph, tt.args.sorting); got != tt.want {
				t.Errorf("IsTopologicallySorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTopSort(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "square",
			args: args{
				graph: [][]int{{1, 2}, {3}, {3}, {}},
			},
		},
		{
			name: "trapezoid",
			args: args{
				graph: [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}},
			},
		},
		{
			name: "long",
			args: args{
				graph: [][]int{{1, 2}, {3}, {3}, {6, 12}, {1, 3, 5}, {7, 8}, {7, 11}, {9, 10}, {7}, {}, {}, {10}, {11}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopSort(tt.args.graph); !IsTopologicallySorted(tt.args.graph, got) {
				t.Errorf("TopSort() = %v, not topologically sorted", got)
			}
		})
	}
}
