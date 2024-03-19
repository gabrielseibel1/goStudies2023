package lc90

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
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
			name: "one",
			args: args{
				nums: []int{1},
			},
			want: [][]int{{}, {1}},
		},
		{
			name: "123",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "122",
			args: args{
				nums: []int{1, 2, 2},
			},
			want: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := subsetsWithDup(tt.args.nums)
			if ok := assert.ElementsMatch(t, tt.want, actual); !ok {
				t.Error("Expected:", tt.want, "Got:", actual)
			}
		})
	}
}
