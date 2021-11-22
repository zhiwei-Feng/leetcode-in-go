package n1060

func missingElement(nums []int, k int) int {
	n := len(nums)
	if k > missingNum(nums, n-1) {
		return nums[n-1] + k - missingNum(nums, n-1)
	}

	low, high := 1, n
	for low < high {
		mid := low + (high-low)/2
		if missingNum(nums, mid) < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low-1] + k - missingNum(nums, low-1)
}

func missingNum(nums []int, idx int) int {
	return nums[idx] - nums[0] - idx
}

func missingElement_On(nums []int, k int) int {
	missings := []int{}
	next := nums[0] + 1
	for i := 1; i < len(nums); i++ {
		for nums[i] != next {
			missings = append(missings, next)
			next++
			if len(missings) >= k {
				return missings[k-1]
			}
		}
		next = nums[i] + 1
	}

	return nums[len(nums)-1] + (k - len(missings))
}
