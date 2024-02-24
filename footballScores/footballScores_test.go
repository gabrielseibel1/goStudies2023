package footballScores

import (
	"reflect"
	"testing"
)

func Test_countMatches(t *testing.T) {
	type args struct {
		teamA []int
		teamB []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "simple",
			args: args{
				teamA: []int{5, 4, 3, 2, 1},
				teamB: []int{3, 4, 5, 6},
			},
			want: []int{3, 4, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countMatches(tt.args.teamA, tt.args.teamB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
