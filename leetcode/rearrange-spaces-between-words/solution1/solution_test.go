package solution1

import (
	"fmt"
	"testing"
)

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		args args
		want string
	}{
		{
			args: args{
				text: "  this   is  a sentence ",
			},
			want: "this   is   a   sentence",
		},
		{
			args: args{
				text: " practice   makes   perfect",
			},
			want: "practice   makes   perfect ",
		},
		{
			args: args{
				text: " practice",
			},
			want: "practice ",
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := reorderSpaces(tt.args.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
