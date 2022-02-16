package n0283

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func moveZeroes_replace(nums []int) {
	left := 0
	for _, v := range nums {
		if v != 0 {
			nums[left] = v
			left++
		}
	}
	for i := left; i < len(nums); i++ {
		nums[i] = 0
	}
}
