package lc130

import (
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "O",
			args: args{
				board: [][]byte{{'O'}},
			},
			want: [][]byte{
				{'O'},
			},
		},
		{
			name: "XXX,XOX,XXX",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X'},
					{'X', 'O', 'X'},
					{'X', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "XXXX,XOOX,XXOX,XOXX",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "XXXXXXX,XXXOXXX,XXOOOXX,XOOOOOX,XXOOOXX,XXXOXXX,XXXXXXX",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'O', 'X', 'X', 'X'},
					{'X', 'X', 'O', 'O', 'O', 'X', 'X'},
					{'X', 'O', 'O', 'O', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'O', 'O', 'X', 'X'},
					{'X', 'X', 'X', 'O', 'X', 'X', 'X'},
					{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
			},
		},
		{
			name: "OOO,OOO,OOO",
			args: args{
				board: [][]byte{
					{'O', 'O', 'O'},
					{'O', 'O', 'O'},
					{'O', 'O', 'O'},
				},
			},
			want: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "OXO,XOX,OXO",
			args: args{
				board: [][]byte{
					{'O', 'X', 'O'},
					{'X', 'O', 'X'},
					{'O', 'X', 'O'},
				},
			},
			want: [][]byte{
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
		},
		{
			name: "big",
			args: args{
				board: [][]byte{
					{'O', 'X', 'X', 'O', 'X'},
					{'X', 'O', 'O', 'X', 'O'},
					{'X', 'O', 'X', 'O', 'X'},
					{'O', 'X', 'O', 'O', 'O'},
					{'X', 'X', 'O', 'X', 'O'},
				},
			},
			want: [][]byte{
				{'O', 'X', 'X', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'O'},
				{'X', 'X', 'X', 'O', 'X'},
				{'O', 'X', 'O', 'O', 'O'},
				{'X', 'X', 'O', 'X', 'O'},
			},
			// {
			//	{'O', 'X', 'X', 'O', 'X'},
			//  {'X', 'X', 'X', 'X', 'X'},
			//  {'O', 'X', 'X', 'X', 'O'},
			//  {'X', 'X', 'O', 'X', 'O'}
			// }
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			if !reflect.DeepEqual(tt.args.board, tt.want) {
				t.Errorf("expected %v, got %v\n", tt.want, tt.args.board)
			}
		})
	}
}
