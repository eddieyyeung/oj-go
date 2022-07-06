package solution2

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
			args{
				s: "babad",
			},
			"bab",
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
		{
			args: args{
				s: "babaddtattarrattatddetartrateedredividerb",
			},
			want: "ddtattarrattatdd",
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
