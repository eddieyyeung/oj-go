package solution2

import (
	"fmt"
	"oj-go/utils"
	"testing"
)

func Test_minTransfers(t *testing.T) {
	type args struct {
		transactions [][]int
	}
	tests := []struct {
		args args
		want int
	}{{
		args: args{
			transactions: [][]int{
				{0, 1, 10},
				{2, 0, 5},
			},
		},
		want: 2,
	},
		{
			args: args{
				transactions: [][]int{
					{0, 1, 10},
					{1, 0, 1},
					{1, 2, 5},
					{2, 0, 5},
				},
			},
			want: 1,
		},
		{
			args: args{
				transactions: [][]int{
					{1, 5, 8},
					{8, 9, 8},
					{2, 3, 9},
					{4, 3, 1},
				},
			},
			want: 4,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got := minTransfers(tt.args.transactions)
			utils.Logger.Sugar().Infow("", "got", got)
			if got != tt.want {
				t.Errorf("minTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
