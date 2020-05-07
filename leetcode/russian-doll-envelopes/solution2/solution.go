// Package solution2 ...
// https://leetcode-cn.com/problems/russian-doll-envelopes/
package solution2

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	sort.Slice(envelopes, func(i, j int) bool {
		e1 := envelopes[i]
		e2 := envelopes[j]
		if e1[0] == e2[0] {
			return e1[1] > e2[1]
		} else {
			return e1[0] < e2[0]
		}
	})
	var arr []int
	for _, v := range envelopes {
		arr = append(arr, v[1])
	}
	return LIS(arr)
}

func LIS(nums []int) int {
	var arr []int
	arr = append(arr, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] > arr[len(arr)-1] {
			arr = append(arr, nums[i])
			continue
		}
		left := 0
		right := len(arr) - 1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] == nums[i] {
				break
			}
			if arr[mid] < nums[i] {
				if mid == len(arr)-1 {
					break
				}
				if nums[i] < arr[mid+1] {
					arr[mid+1] = nums[i]
					break
				}
				left = mid + 1
			}
			if nums[i] < arr[mid] {
				if mid == 0 {
					arr[mid] = nums[i]
					break
				}
				if nums[i] > arr[mid-1] {
					arr[mid] = nums[i]
					break
				}
				right = mid - 1
			}
		}
	}
	return len(arr)
}
