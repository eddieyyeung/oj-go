package findindiceswithindexandvaluedifferenceii

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	maxidx := 0
	minidx := 0
	for r := indexDifference; r < len(nums); r++ {
		l := r - indexDifference
		nl := nums[l]
		if nl > nums[maxidx] {
			maxidx = l
		} else if nl < nums[minidx] {
			minidx = l
		}
		nr := nums[r]
		if nums[maxidx]-nr >= valueDifference {
			return []int{maxidx, r}
		}
		if nr-nums[minidx] >= valueDifference {
			return []int{minidx, r}
		}
	}
	return []int{-1, -1}
}
