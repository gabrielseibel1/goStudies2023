package lc46

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "empty",
			args: args{
				nums: []int{},
			},
			want: [][]int{{}},
		},
		{
			name: "0",
			args: args{
				nums: []int{0},
			},
			want: [][]int{{0}},
		},
		{
			name: "01",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			name: "012",
			args: args{
				nums: []int{0, 1, 2},
			},
			want: [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 1, 0}, {2, 0, 1}},
		},
		{
			name: "5462",
			args: args{
				nums: []int{5, 4, 6, 2},
			},
			want: [][]int{{5, 4, 6, 2}, {5, 4, 2, 6}, {5, 6, 4, 2}, {5, 6, 2, 4}, {5, 2, 4, 6}, {5, 2, 6, 4}, {4, 5, 6, 2}, {4, 5, 2, 6}, {4, 6, 5, 2}, {4, 6, 2, 5}, {4, 2, 5, 6}, {4, 2, 6, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, permute(tt.args.nums))
		})
	}
}
