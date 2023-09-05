package solution2

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		n := nums[i]
		l := i + 1
		r := len(nums) - 1
		for l < r {
			nl := nums[l]
			nr := nums[r]
			s := n + nl + nr

			if s < 0 {
				for l < r {
					if nums[l] != nums[l+1] {
						break
					}
					l++
				}
				l++
			} else if s > 0 {
				for l < r {
					if nums[r] != nums[r-1] {
						break
					}
					r--
				}
				r--
			} else {
				ans = append(ans, []int{n, nl, nr})
				for l < r {
					if nums[l] != nums[l+1] {
						break
					}
					l++
				}
				for l < r {
					if nums[r] != nums[r-1] {
						break
					}
					r--
				}
				l++
				r--
			}
		}
		for i < len(nums)-1 {
			if nums[i] != nums[i+1] {
				break
			}
			i++
		}
	}
	return ans
}
