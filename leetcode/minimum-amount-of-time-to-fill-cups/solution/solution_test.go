package solution

import (
	"fmt"
	"testing"
)

func Test_fillCups(t *testing.T) {
	type args struct {
		amount []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				amount: []int{1, 4, 2},
			},
			want: 4,
		},
		{
			args: args{
				amount: []int{5, 4, 4},
			},
			want: 7,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			if got := fillCups(tt.args.amount); got != tt.want {
				t.Errorf("fillCups() = %v, want %v", got, tt.want)
			}
		})
	}
}
