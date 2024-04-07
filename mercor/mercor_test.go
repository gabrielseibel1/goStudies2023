package main

import "testing"

func TestMatrixSpiral(t *testing.T) {
	type args struct {
		strArr []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0x0",
			args: args{
				strArr: []string{},
			},
			want: "",
		},
		{
			name: "1x1",
			args: args{
				strArr: []string{"[1]"},
			},
			want: "1",
		},
		{
			name: "2x2",
			args: args{
				strArr: []string{"[1,2]", "[3,4]"},
			},
			want: "1,2,4,3",
		},
		{
			name: "5x5",
			args: args{
				// 4 5 6 7 8
				// 1 1 2 2 3
				// 2 3 4 5 6
				// 2 3 4 6 7
				// 1 2 3 4 5
				strArr: []string{"[4,5,6,7,8]", "[1,1,2,2,3]", "[2,3,4,5,6]", "[2,3,4,6,7]", "[1,2,3,4,5]"},
			},
			want: "4,5,6,7,8,3,6,7,5,4,3,2,1,2,2,1,1,2,2,5,6,4,3,3,4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatrixSpiral(tt.args.strArr); got != tt.want {
				t.Errorf("MatrixSpiral() = %v, want %v", got, tt.want)
			}
		})
	}
}
