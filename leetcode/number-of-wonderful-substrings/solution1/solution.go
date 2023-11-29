// 1915. 最美子字符串的数目
// https://leetcode.cn/problems/number-of-wonderful-substrings/description/
package solution

func wonderfulSubstrings(word string) int64 {
	n := len(word)
	var mask int
	var ans int64
	m := map[int]int64{0: 1}
	for i := 0; i < n; i++ {
		mask ^= 1 << int(word[i]-'a')
		// 一个字母出现奇数次的情况
		for j := 0; j < 10; j++ {
			ans += m[mask^(1<<j)]
		}
		// 没有字母出现奇数次的情况
		ans += m[mask]
		m[mask]++
	}
	return ans
}
