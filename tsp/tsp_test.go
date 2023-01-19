package tsp

import (
	"reflect"
	"testing"
)

func Test_combinations(t *testing.T) {
	type args struct {
		on int
		n  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0 on in 2 bits",
			args: args{
				on: 0,
				n:  2,
			},
			want: []int{0},
		},
		{
			name: "2 on in 3 bits",
			args: args{
				on: 2,
				n:  3,
			},
			want: []int{6, 5, 3},
		},
		{
			name: "3 on in 4 bits",
			args: args{
				on: 3,
				n:  4,
			},
			want: []int{14, 13, 11, 7},
		},
		{
			name: "4 on in 4 bits",
			args: args{
				on: 4,
				n:  4,
			},
			want: []int{15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinations(tt.args.on, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinimumCostTour(t *testing.T) {
	type args struct {
		m [][]int
		s int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 []int
	}{
		{
			name: "tsp n=1",
			args: args{
				m: [][]int{
					{5},
				},
				s: 0,
			},
			want:  5, // 5 from 0 to itself
			want1: []int{0, 0},
		},
		{
			name: "tsp n=2",
			args: args{
				m: [][]int{
					{0, 1},
					{1, 0},
				},
				s: 0,
			},
			want:  2, // 1 + 1
			want1: []int{0, 1, 0},
		},
		{
			name: "tsp n=3",
			args: args{
				m: [][]int{
					{0, 1, 2},
					{3, 0, 4},
					{5, 6, 0},
				},
				s: 0,
			},
			want:  10, // 1 + 4 + 5 == 10
			want1: []int{0, 1, 2, 0},
		},
		{
			name: "tsp n=4",
			args: args{
				m: [][]int{
					{0, 4, 1, 9},
					{3, 0, 6, 11},
					{4, 1, 0, 2},
					{6, 5, -4, 0},
				},
				s: 0,
			},
			want:  9, // 9 - 4 + 1 + 3 == 9
			want1: []int{0, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := MinimumCostTour(tt.args.m, tt.args.s)
			if got != tt.want {
				t.Errorf("MinimumCostTour() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("MinimumCostTour() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
