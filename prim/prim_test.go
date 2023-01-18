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
		wantEdges []*graph.Edge
	}{
		{
			name:     "square",
			args:     args{g: graph.Square},
			wantDist: 1,
			wantEdges: []*graph.Edge{
				{
					From: -1,
					To:   0,
					Dist: 0,
				},
				{
					From: 0,
					To:   1,
					Dist: 1,
				},
				{
					From: 1,
					To:   3,
					Dist: 1,
				},
				{
					From: 3,
					To:   2,
					Dist: -1,
				},
			},
		},
		{
			name:     "trapezoid",
			args:     args{g: graph.Trapezoid},
			wantDist: 3,
			wantEdges: []*graph.Edge{
				{
					From: -1,
					To:   0,
					Dist: 0,
				},
				{
					From: 0,
					To:   3,
					Dist: 1,
				},
				{
					From: 3,
					To:   2,
					Dist: -2,
				},
				{
					From: 3,
					To:   1,
					Dist: -1,
				},
				{
					From: 3,
					To:   4,
					Dist: 2,
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
