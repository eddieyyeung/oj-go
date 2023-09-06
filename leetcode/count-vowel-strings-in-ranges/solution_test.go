package countvowelstringsinranges

import (
	"reflect"
	"testing"
)

func Test_vowelStrings(t *testing.T) {
	type args struct {
		words   []string
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				words: []string{"aba", "bcb", "ece", "aa", "e"},
				queries: [][]int{
					{0, 2}, {1, 4}, {1, 1},
				},
			},
			want: []int{2, 3, 0},
		},
		{
			name: "",
			args: args{
				words: []string{"a", "e", "i"},
				queries: [][]int{
					{0, 2},
					{0, 1},
					{2, 2},
				},
			},
			want: []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vowelStrings(tt.args.words, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
