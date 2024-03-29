package island

import "testing"

func TestShortestBridge(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[[0,1],[1,0]]=>1",
			args: args{grid: [][]int{{0, 1}, {1, 0}}},
			want: 1,
		},
		{
			name: "[[0,1,0],[0,0,0],[0,0,1]]=>2",
			args: args{grid: [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}},
			want: 2,
		},
		{
			name: "[[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]=>1",
			args: args{grid: [][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortestBridge(tt.args.grid); got != tt.want {
				t.Errorf("ShortestBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}
