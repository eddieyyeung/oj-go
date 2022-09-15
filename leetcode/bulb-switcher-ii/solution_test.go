package bulb_switcher_ii

import (
	"fmt"
	"testing"
)

func Test_flipLights(t *testing.T) {
	type args struct {
		n       int
		presses int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				n:       1,
				presses: 1,
			},
			want: 2,
		},
		{
			args: args{
				n:       2,
				presses: 1,
			},
			want: 3,
		},
		{
			args: args{
				n:       3,
				presses: 1,
			},
			want: 4,
		},
		{
			args: args{
				n:       3,
				presses: 3,
			},
			want: 8,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := flipLights(tt.args.n, tt.args.presses); got != tt.want {
				t.Errorf("flipLights() = %v, want %v", got, tt.want)
			}
		})
	}
}
