package majorityelement

import "testing"

func TestMajorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,2,3]=3",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},
		{
			name: "[2,2,1,1,1,2,2]=2",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
		{
			name: "[6,5,5]=5",
			args: args{
				nums: []int{6, 5, 5},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MajorityElement(tt.args.nums); got != tt.want {
				t.Errorf("MajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
