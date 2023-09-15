package solution2

var mod = 1337

func pow(x, y, mod int) int {
	ans := 1
	for ; y > 0; y = y / 2 {
		if y&1 != 0 {
			ans = (ans * x) % mod
		}
		x = (x * x) % mod
	}
	return ans
}

func superPow(a int, b []int) int {
	l := len(b)
	var ans int = 1
	for i := 0; i < l; i++ {
		ans = ans * pow(a, b[l-i-1], mod) % mod
		a = pow(a, 10, mod)
	}
	return ans % mod
}
