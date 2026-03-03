# 53. 最大子数组和 — 解法一：动态规划

https://leetcode.cn/problems/maximum-subarray/description/

## 思路

定义 `dp[i]` 为以 `nums[i]` 结尾的子数组的最大和。

**递推关系**：

```
dp[i] = max(nums[i], dp[i-1] + nums[i])
```

- 若 `dp[i-1] < 0`，前面的子数组只会拖累当前元素，从 `nums[i]` 重新开始
- 若 `dp[i-1] >= 0`，拼接前面可以获得更大的和

最终答案为 `max(dp[0..n-1])`。

## 示例演算

输入：`[-2, 1, -3, 4, -1, 2, 1, -5, 4]`，期望输出：`6`

| i | nums[i] | dp[i-1] | dp[i] |
|---|---------|---------|-------|
| 0 | -2      | —       | -2    |
| 1 | 1       | -2      | 1     |
| 2 | -3      | 1       | -2    |
| 3 | 4       | -2      | 4     |
| 4 | -1      | 4       | 3     |
| 5 | 2       | 3       | 5     |
| 6 | 1       | 5       | **6** |
| 7 | -5      | 6       | 1     |
| 8 | 4       | 1       | 5     |

`max(dp) = 6` ✓

## 复杂度

- 时间：O(n)
- 空间：O(n)

## 代码

```go
func maxSubArray(nums []int) int {
    n := len(nums)
    dp := make([]int, n)
    for i := 0; i < n; i++ {
        dp[i] = nums[i]
        if i-1 >= 0 {
            dp[i] = max(dp[i], dp[i-1]+nums[i])
        }
    }
    return slices.Max(dp)
}
```
