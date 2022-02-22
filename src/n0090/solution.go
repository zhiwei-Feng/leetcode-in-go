package n0090

import "sort"

func subsetsWithDup(nums []int) [][]int {
	// 和子集1不同的，对于已经出现过的元素，只能在前一index添加的子集中添加
	sort.Ints(nums)
	ans := [][]int{{}}
	prevAddNum := 0
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] {
			for _, v := range ans[len(ans)-prevAddNum:] {
				tmp := make([]int, len(v))
				copy(tmp, v)
				tmp = append(tmp, nums[i])
				ans = append(ans, tmp)
			}
		} else {
			prevAddNum = len(ans)
			for _, v := range ans {
				tmp := make([]int, len(v))
				copy(tmp, v)
				tmp = append(tmp, nums[i])
				ans = append(ans, tmp)
			}
		}
	}
	return ans
}
