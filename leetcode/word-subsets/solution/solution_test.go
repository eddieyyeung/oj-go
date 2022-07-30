package solution

import (
	"fmt"
	"oj-go/utils"
	"reflect"
	"testing"
)

var (
	logger = utils.NewLogger("")
)

func Test_wordSubsets(t *testing.T) {
	type args struct {
		words1 []string
		words2 []string
	}
	tests := []struct {
		args args
		want []string
	}{
		{
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"e", "o"},
			},
			want: []string{"facebook", "google", "leetcode"},
		},
		{
			args: args{
				words1: []string{"amazon", "apple", "facebook", "google", "leetcode"},
				words2: []string{"lo", "eo"},
			},
			want: []string{"google", "leetcode"},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("case%d", i), func(t *testing.T) {
			got := wordSubsets(tt.args.words1, tt.args.words2)
			logger.Sugar().Infow("", "got", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
