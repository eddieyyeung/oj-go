package solution

import (
	"fmt"
	"testing"
)

func Test_minTransfers(t *testing.T) {
	type args struct {
		transactions [][]int
	}
	tests := []struct {
		args args
		want int
	}{
		{
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
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			if got := minTransfers(tt.args.transactions); got != tt.want {
				t.Errorf("minTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}
