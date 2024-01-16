// 50. Pow(x, n)
// https://leetcode.cn/problems/powx-n/description
package solution

// 解题思路
// 递归 + 位运算
// 假设求 ans = x^77，
// 则 ans = x^(2^5) * x^(2^3) * x^(2^2) * x^(2^1)
// 其中 5 3 2 1 刚好是 77 二进制数 101110 的表示
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / pow(x, -n)
	}
	return pow(x, n)
}

func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var ans float64 = 1

	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n = n >> 1
	}

	return ans
}
