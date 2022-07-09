package solution

import (
	"fmt"
	"oj-go/utils"
	"testing"
)

func Test_minCost(t *testing.T) {
	type args struct {
		costs [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				costs: [][]int{
					{17, 2, 17},
					{16, 16, 5},
					{14, 3, 19},
				},
			},
			want: 10,
		},
		{
			args: args{
				costs: [][]int{
					{7, 6, 2},
				},
			},
			want: 2,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got := minCost(tt.args.costs)
			utils.Logger.Sugar().Infow("", "got", got)
			if got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
