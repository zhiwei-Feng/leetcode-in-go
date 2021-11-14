package n0673

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	count := make([]int, len(nums))
	ans := 1
	number := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 {
					count[i] += count[j]
				}

			}
		}
		if dp[i] > ans {
			ans = dp[i]
			number = count[i]
		} else if dp[i] == ans {
			number += count[i]
		}
	}
	return number
}
