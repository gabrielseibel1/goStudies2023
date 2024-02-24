package lc7

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0",
			args: args{
				x: 0,
			},
			want: 0,
		},
		{
			name: "87",
			args: args{
				x: 87,
			},
			want: 78,
		},
		{
			name: "-87",
			args: args{
				x: -87,
			},
			want: -78,
		},
		{
			name: "-2^31",
			args: args{
				x: -8463847412,
			},
			want: -2147483648,
		},
		{
			name: "-2^31-1",
			args: args{
				x: -9463847412,
			},
			want: 0,
		},
		{
			name: "2^31-1",
			args: args{
				x: 6463847412,
			},
			want: 2147483646,
		},
		{
			name: "2^31+1",
			args: args{
				x: 8463847412,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
