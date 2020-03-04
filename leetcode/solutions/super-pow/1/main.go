package solution

// SuperPow ...
func SuperPow(a int, b []int) int {
	const N = 1337
	remains := []int{}
	remain := a % N
	if remain == 0 {
		return 0
	}
	t := remain
	remains = append(remains, t)
	for {
		t = (t * remain) % N
		if t == remains[0] {
			break
		}
		remains = append(remains, t)
	}
	k := 0
	for i, v := range b {
		if i == len(b)-1 {
			k = (k*10 + v - 1) % len(remains)
		} else {
			k = (k*10 + v) % len(remains)
		}
	}
	return remains[k]
}

// func SuperPow(a int, b []int) int {
// 	const N = 1337
// 	const PhiN = 1140
// 	remain := a % N
// 	if remain == 0 {
// 		return 0
// 	}
// 	k := 0
// 	for _, v := range b {
// 		k = (k*10 + v) % PhiN
// 	}
// 	t := remain
// 	for i := 2; i <= k; i++ {
// 		t = (t * remain) % N
// 	}
// 	return t
// }
