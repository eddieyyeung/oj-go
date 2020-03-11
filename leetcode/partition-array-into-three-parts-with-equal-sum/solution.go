// Package partition_array_into_three_parts_with_equal_sum ...
// https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum/
package partition_array_into_three_parts_with_equal_sum

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		return false
	}
	divide := sum / 3
	i := 0
	t := 0
	for i < len(A) {
		t += A[i]
		if t == divide {
			break
		}
		i++
	}
	if t != divide {
		return false
	}
	j := i + 1
	for j < len(A)-1 {
		t += A[j]
		if t == divide*2 {
			return true
		}
		j++
	}
	return false
}
