package solution

type Item struct {
	Red   int
	Blue  int
	Green int
}

func minCost(costs [][]int) int {
	first := costs[0]
	item := Item{
		Red:   first[0],
		Blue:  first[1],
		Green: first[2],
	}
	for i := 1; i < len(costs); i++ {
		cost := costs[i]
		newItem := Item{
			Red:   cost[0] + min(item.Blue, item.Green),
			Blue:  cost[1] + min(item.Red, item.Green),
			Green: cost[2] + min(item.Red, item.Blue),
		}
		item = newItem
	}
	return min(item.Red, item.Green, item.Blue)
}

func min(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
