package smallest_string_with_a_given_numeric_value

import (
	"testing"
)

func Test_getSmallestString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				n: 3,
				k: 27,
			},
			want: "aay",
		},
		{
			args: args{
				n: 5,
				k: 73,
			},
			want: "aaszz",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := getSmallestString(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("getSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
