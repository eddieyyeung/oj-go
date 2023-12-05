// 2953. 统计完全子字符串
// https://leetcode.cn/problems/count-complete-substrings/description/
package solution

func f(s string, k int, win_size int) int {
	if win_size > len(s) {
		return 0
	}
	var rst, left int
	cnts := make([]int, 26)
	check := func() int {
		for _, cnt := range cnts {
			if cnt > 0 && cnt != k {
				return 0
			}
		}
		return 1
	}
	for right, ch := range s {
		cnts[ch-'a']++
		if right > 0 && abs(int(s[right])-int(s[right-1])) > 2 {
			for left < right {
				cnts[s[left]-'a']--
				left++
			}
		}
		if right-left+1 == win_size {
			rst += check()
			cnts[s[left]-'a']--
			left++
		}
	}
	return rst
}

func countCompleteSubstrings(word string, k int) int {
	var ans int = 0
	for j := 1; j <= 26 && j*k <= len(word); j++ {
		ans += f(word, k, j*k)
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
