package mergesort

import (
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/list"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		head *list.Node
	}
	tests := []struct {
		name string
		args args
		want *list.Node
	}{
		{
			name: "[4,2,1,3]=[1,2,3,4]",
			args: args{
				head: list.New([]int{4, 2, 1, 3}),
			},
			want: list.New([]int{1, 2, 3, 4}),
		},
		{
			name: "[-1,5,3,4,0]=[-1,0,3,4,5]",
			args: args{
				head: list.New([]int{-1, 5, 3, 4, 0}),
			},
			want: list.New([]int{-1, 0, 3, 4, 5}),
		},
		{
			name: "[]=[]",
			args: args{
				head: list.New([]int{}),
			},
			want: list.New([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
