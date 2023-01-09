package jump

import "testing"

func TestJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "23114",
			args: args{
				nums: []int{2, 3, 1, 1, 4},
			},
			want: 2,
		},
		{
			name: "23014",
			args: args{
				nums: []int{2, 3, 0, 1, 4},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Jump(tt.args.nums); got != tt.want {
				t.Errorf("Jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
