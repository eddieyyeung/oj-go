package jianshengziiilcof

var mod int = 1e9 + 7

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

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	pow := n / 3
	remainder := n % 3
	if remainder == 0 {
		return superpow(3, pow, mod)
	} else if remainder == 1 {
		return superpow(3, pow-1, mod) * 4 % mod
	} else {
		return superpow(3, pow, mod) * 2 % mod
	}
}
