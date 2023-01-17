package prim

import (
	"reflect"
	"testing"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func TestMinimumSpanningTree(t *testing.T) {
	type args struct {
		g graph.Graph
	}
	tests := []struct {
		name      string
		args      args
		wantDist  int
		wantEdges []*Edge
	}{
		{
			name:     "square",
			args:     args{g: graph.Square},
			wantDist: 1,
			wantEdges: []*Edge{
				{
					from: -1,
					to:   0,
					dist: 0,
				},
				{
					from: 0,
					to:   1,
					dist: 1,
				},
				{
					from: 1,
					to:   3,
					dist: 1,
				},
				{
					from: 3,
					to:   2,
					dist: -1,
				},
			},
		},
		{
			name:     "trapezoid",
			args:     args{g: graph.Trapezoid},
			wantDist: 3,
			wantEdges: []*Edge{
				{
					from: -1,
					to:   0,
					dist: 0,
				},
				{
					from: 0,
					to:   3,
					dist: 1,
				},
				{
					from: 3,
					to:   2,
					dist: -2,
				},
				{
					from: 3,
					to:   1,
					dist: -1,
				},
				{
					from: 3,
					to:   4,
					dist: 2,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDist, gotEdges := MinimumSpanningTree(tt.args.g)
			if gotDist != tt.wantDist {
				t.Errorf("MinimumSpanningTree() gotDist = %v, want %v", gotDist, tt.wantDist)
			}
			if !reflect.DeepEqual(gotEdges, tt.wantEdges) {
				t.Errorf("MinimumSpanningTree() gotEdges = %v, want %v", gotEdges, tt.wantEdges)
			}
		})
	}
}
