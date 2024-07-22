// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/description/
package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	var ans []int
	var n, m, l int = len(words), len(words[0]), len(words) * len(words[0])
	mp := make(map[string]int)
	for i := 0; i < n; i++ {
		mp[words[i]]++
	}
	for i := 0; i < m && i <= len(s)-l; i++ {
		smp := make(map[string]int)
		for j := 0; j < n; j++ {
			word := s[j*m+i : j*m+i+m]
			smp[word]++
		}
		for j := i; j <= len(s)-l; j += m {
			if j != i {
				smp[s[j-m:j]]--
				smp[s[l-m+j:l+j]]++
			}
			var ok bool = true
			for k, v := range mp {
				if smp[k] != v {
					ok = false
					break
				}
			}
			if ok {
				ans = append(ans, j)
			}
		}
	}
	return ans
}
