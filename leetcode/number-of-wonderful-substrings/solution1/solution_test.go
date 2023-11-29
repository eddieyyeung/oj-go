// 1915. 最美子字符串的数目
// https://leetcode.cn/problems/number-of-wonderful-substrings/description/
package solution

import "testing"

func Test_wonderfulSubstrings(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				word: "aabb",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wonderfulSubstrings(tt.args.word); got != tt.want {
				t.Errorf("wonderfulSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
