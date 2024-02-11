package lc39

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
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
			name: "2367=7",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "235=8",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "2=1",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
		{
			name: "732=18",
			args: args{
				candidates: []int{7, 3, 2},
				target:     18,
			},
			want: [][]int{{7, 7, 2, 2}, {7, 3, 3, 3, 2}, {7, 3, 2, 2, 2, 2}, {3, 3, 3, 3, 3, 3}, {3, 3, 3, 3, 2, 2, 2}, {3, 3, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
