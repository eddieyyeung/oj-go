package solution1

import "math"

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	pow := n / 3
	remainder := n % 3
	if remainder == 0 {
		return int(math.Pow(3, float64(pow)))
	} else if remainder == 1 {
		return int(math.Pow(3, float64(pow-1))) * 2 * 2
	} else {
		return int(math.Pow(3, float64(pow))) * 2
	}
}
