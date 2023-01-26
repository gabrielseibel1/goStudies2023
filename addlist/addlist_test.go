package addlist

import (
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/list"
)

func TestAddTwoNumbers(t *testing.T) {
	type args struct {
		l1 *list.Node
		l2 *list.Node
	}
	tests := []struct {
		name string
		args args
		want *list.Node
	}{
		{
			name: "345+564=807",
			args: args{
				l1: list.New([]int{2, 4, 3}),
				l2: list.New([]int{5, 6, 4}),
			},
			want: list.New([]int{7, 0, 8}),
		},
		{
			name: "0+0=0",
			args: args{
				l1: list.New([]int{0}),
				l2: list.New([]int{0}),
			},
			want: list.New([]int{0}),
		},
		{
			name: "9999999+9999=10009998",
			args: args{
				l1: list.New([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: list.New([]int{9, 9, 9, 9}),
			},
			want: list.New([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name: "9999+9999999=10009998",
			args: args{
				l1: list.New([]int{9, 9, 9, 9}),
				l2: list.New([]int{9, 9, 9, 9, 9, 9, 9}),
			},
			want: list.New([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
