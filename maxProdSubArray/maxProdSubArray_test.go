package maxprodsubarray

import "testing"

func TestMaxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,3,-2,4]=>6",
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			name: "[-2,0,-1]=>0",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProduct(tt.args.nums); got != tt.want {
				t.Errorf("MaxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
