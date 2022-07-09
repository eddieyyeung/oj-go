package solution

import "math"

func reverse(x int) int {
	r := 0
	for x != 0 {
		remain := x % 10
		x = x / 10
		if r > (math.MaxInt32-remain)/10 {
			return 0
		}
		if r < (math.MinInt32-remain)/10 {
			return 0
		}
		r = r*10 + remain
	}
	return r
}
