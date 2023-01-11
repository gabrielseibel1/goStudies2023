package paths

import (
	"reflect"
	"testing"
)

func TestAllPathsSourceTarget(t *testing.T) {
	type args struct {
		graph [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "square",
			args: args{[][]int{{1, 2}, {3}, {3}, {}}},
			want: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			name: "trapezoid",
			args: args{[][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}},
			want: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}},
		},
		{
			name: "complex",
			args: args{[][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}},
			want: [][]int{{0, 1, 7}, {0, 3, 6, 7}, {0, 3, 4, 7}, {0, 1, 4, 7}, {0, 1, 6, 7}, {0, 3, 4, 6, 7}, {0, 1, 4, 6, 7}, {0, 1, 2, 4, 7}, {0, 1, 2, 6, 7}, {0, 1, 5, 6, 7}, {0, 3, 4, 5, 6, 7}, {0, 1, 4, 5, 6, 7}, {0, 1, 2, 4, 6, 7}, {0, 1, 2, 3, 6, 7}, {0, 1, 2, 3, 4, 7}, {0, 1, 2, 4, 5, 6, 7}, {0, 1, 2, 3, 4, 6, 7}, {0, 1, 2, 3, 4, 5, 6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllPathsSourceTarget(tt.args.graph); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPathsSourceTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
