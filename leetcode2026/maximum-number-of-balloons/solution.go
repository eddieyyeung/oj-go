package maximumnumberofballoons

import "math"

func maxNumberOfBalloons(text string) int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}
	balloon := []rune{'b', 'a', 'l', 'o', 'n'}
	min := math.MaxInt
	for _, c := range balloon {
		if c == 'l' || c == 'o' {
			if cnt := count[c] / 2; cnt < min {
				min = cnt
			}
		} else if count[c] < min {
			min = count[c]
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}
