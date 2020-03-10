// Package generateparenthesis ...
// https://leetcode-cn.com/problems/generate-parentheses/
package generateparenthesis

func generateParenthesis(n int) []string {
	re := []string{}
	bt(n, n, 0, "", &re)
	return re
}

func bt(left int, right int, p int, str string, re *[]string) {
	if left == 0 && right == 0 {
		*re = append(*re, str)
		return
	}
	if left > 0 {
		bt(left-1, right, p+1, str+"(", re)
	}
	if right > 0 && p > 0 {
		bt(left, right-1, p-1, str+")", re)
	}
}
