package n0413

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	ans := 0
	d, t := nums[0]-nums[1], 0
	for i := 2; i < len(nums); i++ {
		if nums[i-1]-nums[i] == d {
			t++
		} else {
			t = 0
			d = nums[i-1] - nums[i]
		}
		ans += t
	}
	return ans
}

func numberOfArithmeticSlices_On2(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	ans := 0
	dp := make([][]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]bool, len(nums))
		dp[i][i] = true
		if i < len(nums)-1 {
			dp[i][i+1] = true
		}
	}

	// 处理初始的len=3
	k := 3
	for i := 0; i <= len(nums)-k-1; i++ {
		start, end := i, i+k-1
		if isArithmetic(nums[start : end+1]) {
			ans++
			dp[start][end] = true
		}
	}
	k++
	for ; k <= len(nums); k++ {
		for i := 0; i <= len(nums)-k-1; i++ {
			start, end := i, i+k-1
			if dp[start+1][end-1] && nums[start+1]-nums[start] == nums[start+2]-nums[start+1] && nums[end]-nums[end-1] == nums[end-1]-nums[end-2] {
				ans++
				dp[start][end] = true
			}
		}
	}

	return ans
}

func isArithmetic(nums []int) bool {
	if len(nums) < 3 {
		return true
	}
	n := len(nums)
	D := nums[1] - nums[0]
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] != D {
			return false
		}
	}
	return true
}
