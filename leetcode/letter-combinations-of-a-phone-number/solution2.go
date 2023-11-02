// 17. 电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
package lettercombinationsofaphonenumber

func letterCombinations2(digits string) []string {
	digmap := map[string][]string{
		"1": {},
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	var ans []string
	var dfs func(digits string, str string)
	dfs = func(digits, str string) {
		if len(digits) == 0 {
			if len(str) != 0 {
				ans = append(ans, str)
			}
			return
		}
		digit := digits[:1]
		for _, letter := range digmap[digit] {
			dfs(digits[1:], str+letter)
		}
	}
	dfs(digits, "")
	return ans
}
