package offer0061

import "sort"

func isStraight(nums []int) bool {
	sort.Ints(nums)
	skipChance := 0
	i := 0
	for nums[i] == 0 {
		skipChance++
		i++
	}
	for ; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			continue
		}
		if nums[i]==nums[i+1]{
			return false
		}
		if skipChance >= nums[i+1]-nums[i]-1 {
			skipChance -= nums[i+1] - nums[i] - 1
			continue
		} else {
			return false
		}
	}
	return true
}
