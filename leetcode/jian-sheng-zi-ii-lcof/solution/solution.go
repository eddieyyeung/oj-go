package solution

var mod int = 1e9 + 7

func superpow(a, b, c int) int {
	ans := 1

	for ; b > 0; b = b / 2 {
		if b&1 > 0 {
			ans = (ans * a) % c
		}
		a = (a * a) % c
	}
	return ans
}

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	if r := n % 3; r == 0 {
		return superpow(3, n/3, mod)
	} else if r == 1 {
		return (superpow(3, n/3-1, mod) * 4) % mod
	} else {
		return (superpow(3, n/3, mod) * 2) % mod
	}
}
