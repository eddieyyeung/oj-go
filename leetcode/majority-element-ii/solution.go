// https://leetcode.cn/problems/majority-element-ii/description/
package majorityelementii

// 摩尔投票法, MJRTY - A Fast Majority Vote Algorithm
// https://www.scribd.com/document/319883199/mjrty-ps
func majorityElement(nums []int) []int {
	candidate1, candidate2 := nums[0], nums[0]
	vote1, vote2 := 0, 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if candidate1 == num {
			vote1++
		} else if candidate2 == num {
			vote2++
		} else if vote1 == 0 {
			candidate1 = num
			vote1++
		} else if vote2 == 0 {
			candidate2 = num
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	var ans []int = make([]int, 0, 2)
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if num == candidate1 {
			cnt1++
		}
		if num == candidate2 {
			cnt2++
		}
	}
	if cnt1 > len(nums)/3 {
		ans = append(ans, candidate1)
	}
	if candidate1 != candidate2 && cnt2 > len(nums)/3 {
		ans = append(ans, candidate2)
	}
	return ans
}
