package utils

import (
	"fmt"
	"testing"
)

func TestGenerateNode(t *testing.T) {
	var case1 = []*int{
		PointerInt(1),
		PointerInt(2),
		PointerInt(3),
		PointerInt(4),
		PointerInt(5)}
	node := GenerateNode(case1)
	logger.Sugar().Infow("case1", "node", node)

	var case2 = []*int{
		PointerInt(4),
		PointerInt(5),
		PointerInt(2),
		nil,
		nil,
		PointerInt(3),
		PointerInt(1),
	}
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
