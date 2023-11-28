// 2949. 统计美丽子字符串 II
// https://leetcode.cn/problems/count-beautiful-substrings-ii/description/
package solution

func pSqrt(n int) int {
	var rst int = 1
	for i := 2; i*i <= n; i++ {
		i2 := i * i
		for n%i2 == 0 {
			rst *= i
			n /= i2
		}
		if n%i == 0 {
			rst *= i
			n /= i
		}
	}
	if n > 1 {
		rst *= n
	}
	return rst
	// i := 1
	// for {
	// 	if (i*i)%n == 0 {
	// 		return i
	// 	}
	// 	i++
	// }
}

var mask int = 1065233

func isvowel(c byte) int {
	return 2*(mask>>(c-'a')&1) - 1
}

func beautifulSubstrings(s string, k int) int64 {
	n := len(s)
	k = pSqrt(k * 4)
	var ans int64
	type pair struct{ m, sum int }
	m := map[pair]int64{{0, 0}: 1}
	var sum int
	for i := 0; i < n; i++ {
		sum += isvowel(s[i])
		p := pair{(i + 1) % k, sum}
		ans += m[p]
		m[p]++
	}
	return ans
}
