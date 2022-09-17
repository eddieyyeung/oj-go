package solution

// problem: https://leetcode.cn/problems/largest-substring-between-two-equal-characters/

func maxLengthBetweenEqualCharacters(s string) (ans int) {
	ans = -1

	var letters [26]int
	for i, c := range s {
		if letters[c-'a'] == 0 {
			letters[c-'a'] = i + 1
		} else {
			if i-letters[c-'a'] > ans {
				ans = i - letters[c-'a']
			}
		}
	}
	return
}
