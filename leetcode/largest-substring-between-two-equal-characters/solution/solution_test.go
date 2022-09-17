package solution

import (
	"fmt"
	"testing"
)

func Test_maxLengthBetweenEqualCharacters(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args    args
		wantAns int
	}{
		{
			args: args{
				s: "aa",
			},
			wantAns: 0,
		},
		{
			args: args{
				s: "abca",
			},
			wantAns: 2,
		},
		{
			args: args{
				s: "cbzxy",
			},
			wantAns: -1,
		},
		{
			args: args{
				s: "kvvhywyxzqhbiapzfqiiqzcbqhrdaxiqptqcufymajmuclwyleeahymipmfuupqscrcwo",
			},
			wantAns: 61,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if gotAns := maxLengthBetweenEqualCharacters(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
