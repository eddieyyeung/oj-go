package maximum_rows_covered_by_columns

import (
	"fmt"
	"testing"
)

func Test_maximumRows(t *testing.T) {
	type args struct {
		mat  [][]int
		cols int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{1, 0, 1},
					{0, 1, 1},
					{0, 0, 1},
				},
				cols: 2,
			},
			want: 3,
		},
		{
			args: args{
				mat: [][]int{
					{1},
					{0},
				},
				cols: 1,
			},
			want: 2,
		},
		{
			args: args{
				mat: [][]int{
					{1},
				},
				cols: 1,
			},
			want: 1,
		},
		{
			args: args{
				mat: [][]int{
					{1, 0},
					{0, 1},
					{1, 1},
				},
				cols: 1,
			},
			want: 1,
		},
		{
			args: args{
				mat: [][]int{
					{1, 0, 1, 0, 1, 1, 1},
					{0, 0, 0, 1, 0, 0, 1},
					{1, 0, 0, 0, 0, 1, 0},
				},
				cols: 1,
			},
			want: 0,
		},
		{
			args: args{
				mat: [][]int{
					{1, 0, 1, 1, 1, 1},
					{0, 0, 0, 1, 1, 0},
					{1, 1, 0, 0, 0, 0},
					{0, 0, 1, 1, 0, 1},
				},
				cols: 2,
			},
			want: 1,
		},
		{
			args: args{
				mat: [][]int{
					{1, 0, 1, 0, 1, 0},
				},
				cols: 3,
			},
			want: 1,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := maximumRows(tt.args.mat, tt.args.cols); got != tt.want {
				t.Errorf("maximumRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
