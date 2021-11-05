package n0213

// 为了避免首尾位置同时被取到，我们将合法的盗窃区间划分为两个区间[1:],[:len-1]
// 即第一个区间无法盗窃第一个房子，而第二个区间无法盗窃最后一个房子
func rob(nums []int) int {
	if len(nums) <= 2 {
		return max(nums...)
	}
	return max(_rob(nums[1:]), _rob(nums[:len(nums)-1]))
}

// 不考虑首尾相邻的rob方法
func _rob(nums []int) int {
	if len(nums) <= 2 {
		return max(nums...)
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(x ...int) int {
	maxN := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] > maxN {
			maxN = x[i]
		}
	}
	return maxN
}
