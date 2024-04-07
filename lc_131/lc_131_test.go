package lc131

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "aab",
			args: args{
				s: "aab",
			},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
		{
			name: "aaaaa",
			args: args{
				s: "aaaaa",
			},
			want: [][]string{{"a", "a", "a", "a", "a"}, {"a", "a", "a", "aa"}, {"a", "a", "aa", "a"}, {"a", "a", "aaa"}, {"a", "aa", "a", "a"}, {"a", "aa", "aa"}, {"a", "aaa", "a"}, {"a", "aaaa"}, {"aa", "a", "a", "a"}, {"aa", "a", "aa"}, {"aa", "aa", "a"}, {"aa", "aaa"}, {"aaa", "a", "a"}, {"aaa", "aa"}, {"aaaa", "a"}, {"aaaaa"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
