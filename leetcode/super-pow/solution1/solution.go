package solution1

func superpow(x, y, mod int) int {
	ans := 1
	for ; y > 0; y = y / 2 {
		if y&1 > 0 {
			ans = ((ans % mod) * (x % mod)) % mod
		}
		x = ((x % mod) * (x % mod)) % mod
	}
	return ans
}
