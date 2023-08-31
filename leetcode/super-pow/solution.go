package superpow

func superpow(a, b, n int) int {
	if b == 1 {
		return a % n
	}
	var ans int = 1
	for i := int(0); i < b; i++ {
		ans = ans * a % n
	}
	return ans
}

func superPow(a int, b []int) int {
	var mod int = 1337
	var ans int = 1
	for i := 0; i < len(b); i++ {
		ans = (superpow(ans, 10, mod) * superpow(int(a), int(b[i]), mod)) % mod
	}
	return int(ans)
}
