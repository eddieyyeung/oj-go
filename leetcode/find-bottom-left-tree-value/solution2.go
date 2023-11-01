// 513. 找树左下角的值
// https://leetcode.cn/problems/find-bottom-left-tree-value/description/
package findbottomlefttreevalue

// 把层序遍历改成从右到左遍历，则最后一个节点就是最底层最左边的节点
func findBottomLeftValue2(root *TreeNode) int {
	q := []*TreeNode{root}
	var ans int
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		ans = node.Val
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return ans
}
