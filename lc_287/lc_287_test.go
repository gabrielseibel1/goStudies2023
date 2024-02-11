package lc287

import "testing"

func Test_findDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "13422",
			args: args{
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: "31342",
			args: args{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
		{
			name: "22222",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 2,
		},
		{
			name: "14424",
			args: args{
				nums: []int{1, 4, 4, 2, 4},
			},
			want: 4,
		},
		{
			name: "2222222222",
			args: args{
				nums: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{83, 54, 9, 5, 91, 83, 45, 83, 83, 83, 92, 96, 99, 53, 83, 6, 60, 4, 98, 23, 83, 83, 83, 83, 83, 27, 83, 26, 59, 83, 93, 35, 83, 84, 83, 83, 83, 83, 16, 74, 34, 42, 83, 68, 52, 62, 83, 75, 83, 48, 83, 28, 82, 44, 83, 83, 83, 83, 87, 83, 83, 83, 12, 50, 29, 83, 18, 83, 39, 71, 8, 14, 1, 83, 70, 24, 83, 41, 83, 83, 83, 63, 47, 46, 81, 83, 83, 83, 83, 83, 15, 49, 73, 90, 30, 83, 61, 38, 97, 19},
			},
			want: 83,
		},
		{
			name: "",
			args: args{
				nums: []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
