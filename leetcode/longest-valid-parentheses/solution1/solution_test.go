package solution

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				s: "(())()(()((",
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				s: "())",
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				s: ")()())",
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				s: "(()",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
