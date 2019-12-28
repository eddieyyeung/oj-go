package solution

// ARuleOfDivisibilityBy13 https://www.codewars.com/kata/a-rule-of-divisibility-by-13/go
func ARuleOfDivisibilityBy13(n int) int {
	divisions := []int{1, 10, 9, 12, 3, 4}
	var remainder int
	var nowN int = 0
	var preN int = n
	for nowN != preN {
		sum := 0
		nowN = preN
		tN := nowN
		for i := 0; tN != 0; i++ {
			remainder = tN % 10
			sum += remainder * divisions[i%len(divisions)]
			tN = tN / 10
		}
		preN = sum
	}
	return nowN
}
