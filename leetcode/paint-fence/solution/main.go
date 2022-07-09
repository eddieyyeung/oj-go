package solution

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}
	if n == 2 {
		return k * k
	}
	f1 := 1
	f2 := 1
	m := k - 1
	for i := 2; i < n; i++ {
		ft1 := f2 * m
		f2 = f1 + f2*m
		f1 = ft1
	}

	return (f1 + f2*m) * k
}
