package n0238

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			ans[i] = 1
		} else {
			ans[i] = ans[i-1] * nums[i-1]
		}
	}
	rightProd := 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			rightProd*=nums[i]
			continue
		}
		ans[i] = ans[i] * rightProd
		rightProd*=nums[i]
	}
	return ans
}
