package minimumnumberofarrowstoburstballoons

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	var end int = points[0][1]
	var ans int = 1
	for i := 1; i < len(points); i++ {
		point := points[i]
		if point[0] <= end {
			end = min(end, point[1])
		} else {
			end = point[1]
			ans++
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
