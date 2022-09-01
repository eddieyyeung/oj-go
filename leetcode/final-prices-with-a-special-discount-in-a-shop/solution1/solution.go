package solution1

func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		iv := prices[i]
		for j := i + 1; j < len(prices); j++ {
			jv := prices[j]
			if jv < iv {
				prices[i] -= jv
				break
			}
		}
	}
	return prices
}
