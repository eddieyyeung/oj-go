# 53. 最大子数组和 — 解法二：前缀和

https://leetcode.cn/problems/maximum-subarray/description/

## 思路

子数组 `[l, r]` 的和等于 `presum[r+1] - presum[l]`。

要让差值最大，在从左到右遍历右端点 `r` 时，同步维护已见过的最小前缀和 `min_sum`：

```
ans     = max(ans, sum - min_sum)
min_sum = min(min_sum, sum)
```

其中 `sum` 是当前的前缀和（即 `presum[r+1]`），`min_sum` 初始为 `0`（对应空前缀）。

## 示例演算

输入：`[-2, 1, -3, 4, -1, 2, 1, -5, 4]`，期望输出：`6`

| i | nums[i] | sum | min_sum | ans |
|---|---------|-----|---------|-----|
| — | —       | 0   | 0       | MIN |
| 0 | -2      | -2  | -2      | -2  |
| 1 | 1       | -1  | -2      | 1   |
| 2 | -3      | -4  | -4      | 1   |
| 3 | 4       | 0   | -4      | 4   |
| 4 | -1      | -1  | -4      | 4   |
| 5 | 2       | 1   | -4      | 5   |
| 6 | 1       | 2   | -4      | **6** |
| 7 | -5      | -3  | -4      | 6   |
| 8 | 4       | 1   | -4      | 6   |

`ans = 6` ✓

## 复杂度

- 时间：O(n)
- 空间：O(1)

## 代码

```go
func maxSubArray(nums []int) int {
    sum := 0
    min_sum := 0
    ans := math.MinInt

    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        ans = max(ans, sum-min_sum)
        min_sum = min(min_sum, sum)
    }
    return ans
}
```
