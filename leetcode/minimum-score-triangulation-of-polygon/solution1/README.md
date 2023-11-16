### 题目
[1039. 多边形三角剖分的最低得分](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/description/)

### 解题思路

定义 $dp(i,j)$ 表示区间 $[i,j]$ 之间，多边形三角剖分的最低得分。

定义 $n$ 表示 $values$ 的长度。

在 $dp(i,j)$ 中，枚举 $k \in (i,j)$，将 $[i,j]$ 多边形划分为三个区域：$dfs(i,k)$，三角形 $<i, j, k>$，$dfs(k,j)$，如下图所示：

![](../image_001.png)

则可得到 DP 方程式：
$$
dp(i,j) = \min_{k=i+1}^{j-1}(dfs(i,k)+dfs(k,j)+values[i]*values[j]*values[k]) 
$$

递归边界：$dfs(i,j) = 0 \;when\;j<=i+1 $

递归入口：$dfs(0,n-1)$

时间复杂度：$O(n^3)$

空间复杂度：$O(n^2)$