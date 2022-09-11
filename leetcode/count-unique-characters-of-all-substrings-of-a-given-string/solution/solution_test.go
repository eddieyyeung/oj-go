package solution

import (
	"fmt"
	"testing"
)

func Test_uniqueLetterString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				s: "ABC",
			},
			want: 10,
		},
		{
			args: args{
				s: "ABA",
			},
			want: 8,
		},
		{
			args: args{
				s: "LEETCODE",
			},
			want: 92,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := uniqueLetterString(tt.args.s); got != tt.want {
				t.Errorf("uniqueLetterString() = %v, want %v", got, tt.want)
			}
		})
	}
}
