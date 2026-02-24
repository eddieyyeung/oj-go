package find_the_maximum_factor_score_of_array

import "slices"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func maxScore(nums []int) int64 {
	n := len(nums)
	sufGcd := make([]int, n+1)
	sufLcm := make([]int, n+1)
	sufLcm[n] = 1
	for i, x := range slices.Backward(nums) {
		sufGcd[i] = gcd(sufGcd[i+1], x)
		sufLcm[i] = lcm(sufLcm[i+1], x)
	}
	ans := sufGcd[0] * sufLcm[0]
	preGcd, preLcm := 0, 1
	for i, x := range nums {
		ans = max(ans, gcd(preGcd, sufGcd[i+1])*lcm(preLcm, sufLcm[i+1]))
		preGcd = gcd(preGcd, x)
		preLcm = lcm(preLcm, x)
	}
	return int64(ans)
}
