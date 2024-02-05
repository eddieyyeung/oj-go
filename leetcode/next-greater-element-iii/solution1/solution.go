package solution

import (
	"math"
	"strconv"
)

func nextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))
	i := len(nums) - 2
	// 1. 首先从后向前查找第一个顺序对 (i,i+1)，满足 a[i]<a[i+1]。
	//    这样「较小数」即为 a[i]。此时 [i+1,n) 必然是下降序列。
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	// 2. 如果找到了顺序对，那么在区间 [i+1,n) 中从后向前查找第一个元素 j 满足 a[i]<a[j]。
	//    这样「较大数」即为 a[j]。交换 a[i] 与 a[j]
	for j := len(nums) - 1; j > i; j-- {
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j]
			break
		}
	}
	// 3. 此时可以证明区间 [i+1,n) 必为降序。
	//    我们可以直接使用双指针反转区间 [i+1,n) 使其变为升序，而无需对该区间进行排序。
	reverse(nums[i+1:])
	ans, _ := strconv.Atoi(string(nums))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}

func reverse(arr []byte) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
