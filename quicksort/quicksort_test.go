package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[3,2,1,4,5]",
			args: args{
				s: []int{3, 2, 1, 4, 5},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[5,4,2,1,3]",
			args: args{
				s: []int{5, 4, 2, 1, 3},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "[1]",
			args: args{
				s: []int{1},
			},
			want: []int{1},
		},
		{
			name: "[]",
			args: args{
				s: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
