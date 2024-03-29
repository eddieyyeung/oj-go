// 1143. 最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/description/
package solution

import "testing"

func Test_longestCommonSubsequence(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				text1: "abcba",
				text2: "abcbcba",
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				text1: "abcde",
				text2: "acebcde",
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				text1: "abcde",
				text2: "ace",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				text1: "abc",
				text2: "abc",
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				text1: "abc",
				text2: "def",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
