package bellmanford

import "testing"

func TestMaxProbability(t *testing.T) {
	type args struct {
		n        int
		edges    [][]int
		succProb []float64
		start    int
		end      int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2",
			args: args{
				n:        3,
				edges:    [][]int{{0, 1}, {1, 2}, {0, 2}},
				succProb: []float64{0.5, 0.5, 0.2},
				start:    0,
				end:      2,
			},
			want: 0.25,
		},
		{
			name: "n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2",
			args: args{
				n:        3,
				edges:    [][]int{{0, 1}, {1, 2}, {0, 2}},
				succProb: []float64{0.5, 0.5, 0.3},
				start:    0,
				end:      2,
			},
			want: 0.3,
		},
		{
			name: "n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2",
			args: args{
				n:        3,
				edges:    [][]int{{0, 1}},
				succProb: []float64{0.5},
				start:    0,
				end:      2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProbability(tt.args.n, tt.args.edges, tt.args.succProb, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("MaxProbability() = %v, want %v", got, tt.want)
			}
		})
	}
}
