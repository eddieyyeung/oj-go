// 2246. 相邻字符不同的最长路径
// https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/description/
package solution

import "testing"

func Test_longestPath(t *testing.T) {
	type args struct {
		parent []int
		s      string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				parent: []int{-1, 0, 0, 1, 1, 3, 3, 3},
				s:      "aaabcbbb",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				parent: []int{-1, 0, 0, 1, 1, 2},
				s:      "abacbe",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				parent: []int{-1, 0, 0, 0},
				s:      "aabc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPath(tt.args.parent, tt.args.s); got != tt.want {
				t.Errorf("longestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
