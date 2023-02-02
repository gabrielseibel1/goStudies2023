package bipartite

import (
	"testing"

	"github.com/gabrielseibel1/goStudies2023/graph"
)

func TestIsBipartite(t *testing.T) {
	type args struct {
		g graph.Graph
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Square=>true",
			args: args{g: graph.Square},
			want: true,
		},
		{
			name: "Trapezoid=>false",
			args: args{g: graph.Trapezoid},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBipartite(tt.args.g); got != tt.want {
				t.Errorf("IsBipartite() = %v, want %v", got, tt.want)
			}
		})
	}
}
