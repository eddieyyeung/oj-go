package solution3

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				s: "abb",
			},
			want: "bb",
		},
		{
			args: args{
				s: "aacabdkacaa",
			},
			want: "aca",
		},
		{
			args: args{
				s: "babaddtattarrattatddetartrateedredividerb",
			},
			want: "ddtattarrattatdd",
		},
		{
			args{
				s: "babad",
			},
			"aba",
		},
		{
			args{
				s: "cbbd",
			},
			"bb",
		},
		{
			args: args{s: "ac"},
			want: "a",
		},
		{
			args: args{
				s: "adam",
			},
			want: "ada",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			got := longestPalindrome(tt.args.s)
			t.Log(got)
			if got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
