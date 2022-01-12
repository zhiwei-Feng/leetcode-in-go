package n0334

func increasingTriplet(nums []int) bool {
	n := len(nums)
	leftMin := make([]int, n)
	rightMax := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			leftMin[i] = nums[0]
			continue
		}

		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			rightMax[i] = nums[n-1]
			continue
		}

		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] {
			return true
		}
	}

	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
