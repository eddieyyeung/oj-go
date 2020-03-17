// Package longest_common_prefix ...
// https://leetcode-cn.com/problems/longest-common-prefix/
package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	str := strs[0]
	for i := 1; i < len(strs); i++ {
		j := 0
		for j < len(str) {
			if j >= len(strs[i]) {
				break
			}
			if str[j] != strs[i][j] {
				break
			}
			j++
		}
		str = str[:j]
	}
	return str
}
