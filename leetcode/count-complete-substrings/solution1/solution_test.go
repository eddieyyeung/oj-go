package solution

import "testing"

func Test_countCompleteSubstrings(t *testing.T) {
	type args struct {
		word string
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				word: "gvgvvgv",
				k:    2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				word: "aaa",
				k:    1,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				word: "igigee",
				k:    2,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				word: "aaabbbccc",
				k:    3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompleteSubstrings(tt.args.word, tt.args.k); got != tt.want {
				t.Errorf("countCompleteSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
