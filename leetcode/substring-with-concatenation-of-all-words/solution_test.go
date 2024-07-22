package substringwithconcatenationofallwords

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				s:     "a",
				words: []string{"a", "a"},
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				s:     "lingmindraboofooowingdingbarrwingmonkeypoundcake",
				words: []string{"fooo", "barr", "wing", "ding", "wing"},
			},
			want: []int{13},
		},
		{
			name: "",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			want: []int{0, 9},
		},
		{
			name: "",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: []int{},
		},
		{
			name: "",
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			want: []int{6, 9, 12},
		},
		{
			name: "",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "good"},
			},
			want: []int{8},
		},
		{
			name: "",
			args: args{
				s:     "aaaaaaaaaaaa",
				words: []string{"aaa", "aaa", "aaa"},
			},
			want: []int{0, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
