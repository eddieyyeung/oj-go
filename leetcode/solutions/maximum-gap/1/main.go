package solution

import (
	"strconv"
)

func countSort(arr []int, exp int) {
	output := make([]int, len(arr))
	count := [10]int{}
	for _, v := range arr {
		count[(v/exp)%10]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		t := (v / exp) % 10
		count[t]--
		output[count[t]] = v
	}
	for i, v := range output {
		arr[i] = v
	}
}

func getMax(arr []int) (index int, value int) {
	maxI := 0
	maxV := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maxV {
			maxI, maxV = i, arr[i]
		}
	}
	return maxI, maxV
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	_, maxV := getMax(nums)
	maxLen := len(strconv.Itoa(maxV))
	exp := 1
	for i := 0; i < maxLen; i++ {
		countSort(nums, exp)
		exp *= 10
	}
	gap := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if nGap := nums[i] - nums[i-1]; nGap > gap {
			gap = nGap
		}
	}
	return gap
}

// Run ...
func Run(nums []int) int {
	return maximumGap(nums)
}
