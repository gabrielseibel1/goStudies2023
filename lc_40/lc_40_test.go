package lc40

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum2(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example1",
			args: args{candidates: []int{10, 1, 2, 7, 6, 1, 5}, target: 8},
			want: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name: "example2",
			args: args{candidates: []int{2, 5, 2, 1, 2}, target: 5},
			want: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			name: "test172",
			args: args{candidates: []int{2, 5, 1, 1, 2, 3, 3, 3, 1, 2, 2}, target: 5},
			want: [][]int{{1, 1, 1, 2}, {1, 1, 3}, {1, 2, 2}, {2, 3}, {5}},
		},
		{
			name: "MLE",
			args: args{
				candidates: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				target:     30,
			},
			want: [][]int{{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum2(tt.args.candidates, tt.args.target); !assert.ElementsMatch(t, tt.want, got) {
				t.Errorf("%s : combinationSum2() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
