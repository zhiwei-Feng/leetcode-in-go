package meituan0002

import "math"

// 给一个数组，可以对任意子数组翻转一次，也可以不翻，求最大连续和

func solve(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	curMax := 0
	for i := 0; i < n; i++ {
		curMax = max(0, curMax+nums[i])
		if i == 0 {
			dp1[i] = curMax
		} else {
			dp1[i] = max(dp1[i-1], curMax)
		}
	}
	for j := n - 1; j >= 0; j-- {
		if j==n-1{
			dp2[j] = nums[j]
		}else{
			dp2[j] = max(dp2[j+1]+nums[j], nums[j])
		}
	}
	ans := math.MinInt32
	for i := 0; i < n-1; i++ {
		ans = max(ans, dp1[i]+dp2[i+1])
	}
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
