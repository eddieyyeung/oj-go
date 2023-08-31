package bashuzupaichengzuixiaodeshulcof

import (
	"fmt"
	"sort"
	"strconv"
)

func intJoin(nums []int) string {
	var s string
	for _, n := range nums {
		s += strconv.Itoa(n)
	}
	return s
}

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		n1 := fmt.Sprintf("%d%d", nums[i], nums[j])
		n2 := fmt.Sprintf("%d%d", nums[j], nums[i])
		return n1 < n2
	})
	return intJoin(nums)
}
