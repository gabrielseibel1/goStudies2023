package dijkstra

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	type args struct {
		times [][]int
		n     int
		k     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2 |= 2",
			args: args{times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, n: 4, k: 2},
			want: 2,
		},
		{
			name: "times = [[1,2,1]], n = 2, k = 1 |= 1",
			args: args{times: [][]int{{1, 2, 1}}, n: 2, k: 1},
			want: 1,
		},
		{
			name: "times = [[1,2,1]], n = 2, k = 2 |= -1",
			args: args{times: [][]int{{1, 2, 1}}, n: 2, k: 2},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NetworkDelayTime(tt.args.times, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("NetworkDelayTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
