// 2951. 找出峰值
// https://leetcode.cn/problems/find-the-peaks/description/
package solution

func findPeaks(mountain []int) []int {
	var ans []int
	for i := 1; i < len(mountain)-1; i++ {
		if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
			ans = append(ans, i)
		}
	}
	return ans
}
