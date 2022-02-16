package n0053

func maxSubArray(nums []int) int {
	ans := 0
	tmpMax := 0
	for i, v := range nums {
		if i == 0 {
			tmpMax = v
		} else {
			tmpMax = max(nums[i], tmpMax+nums[i])
		}
		if tmpMax > ans {
			ans = tmpMax
		}
	}
	return ans
}

// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	ans := dp[0]
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(nums[i], dp[i-1]+nums[i])
// 		if dp[i] > ans {
// 			ans = dp[i]
// 		}
// 	}

// 	return ans
// }

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
