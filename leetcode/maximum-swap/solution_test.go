package maximum_swap

import (
	"fmt"
	"testing"
)

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		args    args
		wantAns int
	}{
		{
			args: args{
				num: 2736,
			},
			wantAns: 7236,
		},
		{
			args: args{
				num: 9973,
			},
			wantAns: 9973,
		},
		{
			args: args{
				num: 0,
			},
			wantAns: 0,
		},
		{
			args: args{
				num: 1993,
			},
			wantAns: 9913,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if gotAns := maximumSwap(tt.args.num); gotAns != tt.wantAns {
				t.Errorf("maximumSwap() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
