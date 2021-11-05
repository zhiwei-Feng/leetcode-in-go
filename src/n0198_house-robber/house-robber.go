package n0198

func rob(nums []int) int {
    if len(nums)<=2 {
        return max(nums...)
    }
    dp := make([]int, len(nums))
    dp[0]=nums[0]
    dp[1]=max(nums[0],nums[1])
    for i:=2;i<len(nums);i++{
        dp[i]=max(dp[i-2]+nums[i], dp[i-1])
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