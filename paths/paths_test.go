package paths

import (
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func TestAllPathsSourceTarget(t *testing.T) {
	type args struct {
		graph graph.Graph
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "square",
			args: args{graph.Square},
			want: [][]int{{0, 1, 3}, {0, 2, 3}},
		},
		{
			name: "trapezoid",
			args: args{graph.Trapezoid},
			want: [][]int{{0, 4}, {0, 3, 4}, {0, 1, 3, 4}, {0, 1, 2, 3, 4}},
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
