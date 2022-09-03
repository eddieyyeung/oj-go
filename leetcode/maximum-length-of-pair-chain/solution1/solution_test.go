package solution1

import (
	"fmt"
	"testing"
)

func Test_findLongestChain(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				pairs: [][]int{
					{1, 2},
					{2, 3},
					{3, 4},
				},
			},
			want: 2,
		},
		{
			args: args{
				pairs: [][]int{
					{-10, -8},
					{8, 9},
					{-5, 0},
					{6, 10},
					{-6, -4},
					{1, 7},
					{9, 10},
					{-4, 7},
				},
			},
			want: 4,
		},
		{
			args: args{
				pairs: [][]int{
					{-10, -8},
					{8, 9},
					{-5, 0},
					{6, 10},
					{-6, -4},
					{1, 7},
					{8, 10},
					{-4, 7},
				},
			},
			want: 4,
		},
		{
			args: args{
				pairs: [][]int{
					{9, 10},
					{-9, 9},
					{-6, 1},
					{-4, 1},
					{8, 10},
					{7, 10},
					{9, 10},
					{2, 10},
				},
			},
			want: 2,
		},
		{
			args: args{
				pairs: [][]int{
					{-4, -2},
					{9, 10},
					{-2, 1},
					{6, 9},
					{7, 9},
					{-10, -3},
					{-8, 5},
					{5, 8},
				},
			},
			want: 4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := findLongestChain(tt.args.pairs); got != tt.want {
				t.Errorf("findLongestChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
