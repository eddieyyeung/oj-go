package solution3

import (
	"testing"
)

func Test_greatestLetter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				s: "lEeTcOdE",
			},
			want: "E",
		},
		{
			args: args{
				s: "arRAzFif",
			},
			want: "R",
		},
		{
			args: args{
				s: "AbCdEfGhIjK",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := greatestLetter(tt.args.s); got != tt.want {
				t.Errorf("greatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
