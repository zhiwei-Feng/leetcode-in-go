package netease0001

// 将指定的元素1、6、3移动到末尾，并且需要保持所有元素的相对顺序。
func solve(nums []int) []int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if !is163(nums[i]) {
			tmp := nums[i]
			for j := i - 1; j >= idx; j-- {
				nums[j+1] = nums[j]
			}
			nums[idx] = tmp
			idx++
		}
	}
	return nums
}

func is163(x int) bool {
	return x == 1 || x == 6 || x == 3
}
