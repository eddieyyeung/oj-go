// 372. 超级次方
// https://leetcode.cn/problems/super-pow/description/
package solution

const mod = 1337

// x^n
// n = 13
// x^13 = (x^6 * x^6) * x = (x^6)^2 * x
//
//					= ((x^3)^2)^2 * x
//					= ((x^2 * x)^2)^2 * x		// 可以用递归实现的
//		  		= x^2^2^2 * x^2^2 * x
//					= x^(2^3) * x^(2^2) * x^(2^0)  n = 13 二进制的表达 1101
//	 0. t = x
//	    ans = t
//	 1. t = x * x
//	 2. t = t * t = (x*x) * (x*x)
//	    ans = ans * t
func pow(x, n, m int) int {
	var ans int = 1
	for n > 0 {
		if n&1 == 1 {
			ans = ans * x % m
		}
		n = n >> 1
		x = x * x % m
	}
	return ans
}

func superPow(a int, b []int) int {
	var ans int = 1
	for _, e := range b {
		ans = pow(ans, 10, mod) * pow(a, e, mod) % mod
	}
	return ans
}
