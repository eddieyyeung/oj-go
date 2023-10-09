package besttimetobuyandsellstock

func maxProfit(prices []int) int {
	var buy, sell int
	buy = prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		if price-buy > sell {
			sell = price - buy
		}
		if price < buy {
			buy = price
		}
	}
	return sell
}
