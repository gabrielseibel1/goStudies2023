package bridges

import (
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func TestFindBridges(t *testing.T) {
	type args struct {
		g graph.Graph
	}
	tests := []struct {
		name string
		args args
		want []*graph.Edge
	}{
		{
			name: "square",
			args: args{g: graph.Square},
			want: []*graph.Edge{},
		},
		{
			name: "trapezoid",
			args: args{g: graph.Trapezoid},
			want: []*graph.Edge{},
		},
		{
			name: "no-bridge",
			args: args{g: graph.NoBridge},
			want: []*graph.Edge{},
		},
		{
			name: "bridge",
			args: args{g: graph.Bridge},
			want: []*graph.Edge{
				{
					From: 4,
					To:   5,
					Dist: 1,
				},
				{
					From: 3,
					To:   4,
					Dist: 1,
				},
				{
					From: 2,
					To:   3,
					Dist: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBridges(tt.args.g); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBridges() = %v, want %v", got, tt.want)
			}
		})
	}
}
