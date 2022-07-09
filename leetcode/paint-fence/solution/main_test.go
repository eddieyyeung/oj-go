package solution

import (
	"fmt"
	"oj-go/utils"
	"testing"
)

func Test_numWays(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n: 3,
				k: 2,
			},
			want: 6,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got := numWays(tt.args.n, tt.args.k)
			utils.Logger.Sugar().Infow("", "got", got)
			if got != tt.want {
				t.Errorf("numWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
