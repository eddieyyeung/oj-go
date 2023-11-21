// 968. 监控二叉树
// https://leetcode.cn/problems/binary-tree-cameras/description/
package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分类讨论
// cover 安装摄像头
// cover_fa 不安装摄像头，且父节点安装摄像头
// cover_ch 不安装摄像头，且至少一个子节点安装摄像头
// 则：
// cover = min(l_cover,l_cover_fa,l_cover_ch) + min(r_cover,r_cover_fa,r_cover_ch) + 1
// cover_fa = min(l_cover,l_cover_ch) + min(r_cover,r_cover_ch)
// cover_ch = min(l_cover+r_cover_ch,l_cover_ch+r_cover,l_cover+r_cover)
// 简化为
// cover_fa = min(l_cover,l_cover_ch) + min(r_cover,r_cover_ch)
// cover_ch = cover_fa + max(0, min(l_cover-l_cover_ch,r_cover-r_cover_ch))
// cover = min(l_cover,l_cover_fa) + min(r_cover,r_cover_fa) + 1
func minCameraCover(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int, int)

	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return 1e9, 0, 0
		}
		l_cover, l_cover_fa, l_cover_ch := dfs(node.Left)
		r_cover, r_cover_fa, r_cover_ch := dfs(node.Right)
		cover_fa := min(l_cover, l_cover_ch) + min(r_cover, r_cover_ch)
		cover_ch := cover_fa + max(0, min(l_cover-l_cover_ch, r_cover-r_cover_ch))
		cover := min(l_cover, l_cover_fa) + min(r_cover, r_cover_fa) + 1
		return cover, cover_fa, cover_ch
	}
	cover, _, cover_ch := dfs(root)
	return min(cover, cover_ch)
}

func max(nums ...int) int {
	var ans int = math.MinInt
	for _, num := range nums {
		if num > ans {
			ans = num
		}
	}
	return ans
}

func min(nums ...int) int {
	var ans int = math.MaxInt
	for _, num := range nums {
		if num < ans {
			ans = num
		}
	}
	return ans
}
