package solution

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				123,
			},
			want: 321,
		},
		{
			args: args{x: -123},
			want: -321,
		},
		{
			args: args{x: 120},
			want: 21,
		},
		{
			args: args{x: 0},
			want: 0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got := reverse(tt.args.x)
			t.Log("got", got)
			if got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
