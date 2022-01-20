package offer0057

func twoSum(nums []int, target int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if target == nums[i]+nums[j] {
			return []int{nums[i], nums[j]}
		}

		if target > nums[i]+nums[j] {
			i++
		} else {
			j--
		}
	}
	return []int{}
}
