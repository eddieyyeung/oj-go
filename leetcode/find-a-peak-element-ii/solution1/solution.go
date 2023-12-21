// 1901. 寻找峰值 II
// https://leetcode.cn/problems/find-a-peak-element-ii/description/
package solution

func findPeakGrid(mat [][]int) []int {
	left, right := 0, len(mat)-2
	for left <= right {
		mid := left + (right-left)/2
		max_idx := indexOfMax(mat[mid])
		if mat[mid][max_idx] < mat[mid+1][max_idx] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return []int{left, indexOfMax(mat[left])}
}

func indexOfMax(a []int) (idx int) {
	for i, x := range a {
		if x > a[idx] {
			idx = i
		}
	}
	return
}
