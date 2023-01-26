package utils

import (
	"fmt"
	"testing"
)

func TestGenerateNode(t *testing.T) {
	var case1 = []int{1, 2, 3, 4, 5}
	node := GenerateNode(case1)
	logger.Sugar().Infow("case1", "node", node)

	var case2 = []int{4, 5, 2, -1, -1, 3, 1}
	node = GenerateNode(case2)
	logger.Sugar().Infow("case2", "node", node)
}

func TestCount1(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				x: 7,
			},
			want: 3,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			if got := Count1(tt.args.x); got != tt.want {
				t.Errorf("Count1() = %v, want %v", got, tt.want)
			}
		})
	}
}
