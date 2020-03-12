// Package greatest_common_divisor_of_strings ...
// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings/
package greatest_common_divisor_of_strings

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) == len(str2) && str1 == str2 {
		return str1
	}
	if len(str1) > len(str2) && str1[:len(str2)] == str2 {
		return gcdOfStrings(str1[len(str2):], str2)
	}
	if len(str2) > len(str1) && str2[:len(str1)] == str1 {
		return gcdOfStrings(str2[len(str1):], str1)
	}
	return ""
}
