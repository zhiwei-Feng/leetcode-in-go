package offerii0008

func minSubArrayLen(target int, nums []int) int {
	winS := 0
	winSum := 0
	ans := len(nums) + 1
	for winE := 0; winE < len(nums); winE++ {
		winSum += nums[winE]
		for winSum >= target {
			ans = min(ans, winE-winS+1)
			winSum -= nums[winS]
			winS++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
