// 76. 最小覆盖子串
// https://leetcode.cn/problems/minimum-window-substring/description/
package solution

func minWindow(s string, t string) string {
	mp_t := make(map[byte]int)
	var valid int
	for _, ch := range t {
		mp_t[byte(ch)]++
	}
	mp_s := make(map[byte]int)
	left := 0
	var start, end int = 0, 1e9
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if mp_t[ch] != 0 {
			mp_s[ch]++
			if mp_s[ch] == mp_t[ch] {
				valid++
			}
		}
		for valid == len(mp_t) {
			if i+1-left < end-start {
				start, end = left, i+1
			}
			if mp_t[s[left]] != 0 {
				if mp_t[s[left]] == mp_s[s[left]] {
					valid--
				}
				mp_s[s[left]]--
			}
			left++
		}
	}
	if end == 1e9 {
		return ""
	}
	return s[start:end]
}
