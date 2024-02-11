package lc78

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nothing",
			args: args{},
			want: [][]int{{}},
		},
		{
			name: "0",
			args: args{
				nums: []int{0},
			},
			want: [][]int{{}, {0}},
		},
		{
			name: "01",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{}, {0}, {1}, {0, 1}},
		},
		{
			name: "012",
			args: args{
				nums: []int{0, 1, 2},
			},
			want: [][]int{{}, {0}, {1}, {2}, {0, 1}, {0, 2}, {1, 2}, {0, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestCombinations(t *testing.T) {
	type args struct {
		slice []int
		size  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nothing",
			args: args{},
			want: [][]int{},
		},
		{
			name: "0-0",
			args: args{
				slice: []int{0},
				size:  0,
			},
			want: [][]int{},
		},
		{
			name: "0-1",
			args: args{
				slice: []int{0},
				size:  1,
			},
			want: [][]int{{0}},
		},
		{
			name: "01-0",
			args: args{
				slice: []int{0, 1},
				size:  0,
			},
			want: [][]int{},
		},
		{
			name: "01-1",
			args: args{
				slice: []int{0, 1},
				size:  1,
			},
			want: [][]int{{0}, {1}},
		},
		{
			name: "01-2",
			args: args{
				slice: []int{0, 1},
				size:  2,
			},
			want: [][]int{{0, 1}},
		},
		{
			name: "012-0",
			args: args{
				slice: []int{0, 1, 2},
				size:  0,
			},
			want: [][]int{},
		},
		{
			name: "012-1",
			args: args{
				slice: []int{0, 1, 2},
				size:  1,
			},
			want: [][]int{{0}, {1}, {2}},
		},
		{
			name: "012-2",
			args: args{
				slice: []int{0, 1, 2},
				size:  2,
			},
			want: [][]int{{0, 1}, {0, 2}, {1, 2}},
		},
		{
			name: "012-3",
			args: args{
				slice: []int{0, 1, 2},
				size:  3,
			},
			want: [][]int{{0, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinations(tt.args.slice, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}
