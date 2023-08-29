package superpow

func superpow(a, b, n int64) int64 {
	if b == 1 {
		return a % n
	}
	var ans int64 = 1
	for i := int64(0); i < b; i++ {
		ans = ((ans % n) * (a % n)) % n
	}
	return ans
}

func superPow(a int, b []int) int {
	var mod int64 = 1337
	var ans int64 = 1
	for i := 0; i < len(b); i++ {
		ans = (superpow(ans, 10, mod) * superpow(int64(a), int64(b[i]), mod)) % mod
	}
	return int(ans)
}
