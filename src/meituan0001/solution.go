package meituan0001

import "sort"

// 给一个带有重复数字的数组，找出最长的不连续的序列的长度（好像是上升的）。
// 输入：1,2,3,5,6,7
// 输出：4，最长为【1,3,5,7】
func solve(nums []int) int {
	sort.Ints(nums)
	nums = unique(nums)
	dp := make([]int, len(nums)+1) // 表示到i位置的最长不连续序列长度
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= len(nums); i++ {
		dp[i] = dp[i-1] // 不选
		if nums[i-1] > nums[i-2]+1 {
			dp[i] = max(dp[i], dp[i-1]+1)
		}
		dp[i] = max(dp[i], dp[i-2]+1) // 很关键, 考虑1,3,4,5的情况下,当遍历到5的时候仅看前一位是不行的
	}
	return dp[len(nums)]
}

func unique(nums []int) []int {
	idx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return nums[:idx]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
