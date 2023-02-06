package gold

import "testing"

func TestGetMaximumGold(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "[[0,6,0],[5,8,7],[0,9,0]]=>24",
			args: args{grid: [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}},
			want: 24,
		},
		{
			name: "[[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]",
			args: args{grid: [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}},
			want: 28,
		},
		{
			name: "[[0,0,2,0],[13,20,36,0],[0,31,27,0]]=>129",
			args: args{grid: [][]int{{0, 0, 2, 0}, {13, 20, 36, 0}, {0, 31, 27, 0}}},
			want: 129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMaximumGold(tt.args.grid); got != tt.want {
				t.Errorf("GetMaximumGold() = %v, want %v", got, tt.want)
			}
		})
	}
}
