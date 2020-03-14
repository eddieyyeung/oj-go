// Package longest_increasing_subsequence ...
// https://leetcode-cn.com/problems/longest-increasing-subsequence/
package longest_increasing_subsequence

// 贪心 + 二分查找
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := []int{}
	dp = append(dp, nums[0])
Loop:
	for i := 1; i < len(nums); i++ {
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
			continue
		}
		l := 0
		r := len(dp) - 1
		for l <= r {
			mid := (l + r) / 2
			if dp[mid] == nums[i] {
				dp[mid] = nums[i]
				continue Loop
			} else if dp[mid] < nums[i] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		dp[r+1] = nums[i]
	}
	return len(dp)
}

// Solution: 动态规划
// func lengthOfLIS(nums []int) int {
// 	if len(nums) < 1 {
// 		return 0
// 	}
// 	dp := make([]int, len(nums))
// 	dp[0] = 1
// 	maxL := 1
// 	for i := 1; i < len(nums); i++ {
// 		ml := 1
// 		for j := 0; j < i; j++ {
// 			if nums[j] < nums[i] && dp[j]+1 > ml {
// 				ml = dp[j] + 1
// 			}
// 		}
// 		dp[i] = ml
// 		if ml > maxL {
// 			maxL = ml
// 		}
// 	}
// 	return maxL
// }

// Solution1: 回溯
// func lengthOfLIS(nums []int) int {
// 	maxl := 0
// 	for i := 0; i < len(nums); i++ {
// 		bc(nums, i, 1, &maxl)
// 	}
// 	return maxl
// }
//
// func bc(nums []int, idx int, l int, maxl *int) {
// 	if l > *maxl {
// 		*maxl = l
// 	}
// 	for i := idx + 1; i < len(nums); i++ {
// 		if nums[i] > nums[idx] {
// 			bc(nums, i, l+1, maxl)
// 		}
// 	}
// }
