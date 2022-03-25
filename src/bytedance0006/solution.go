package bytedance0006

// 寻找旋转排序数组的最大值
func solve(nums []int) int {
	l , r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)/2
		if nums[mid]>nums[l] {
			l = mid
		}else {
			r = mid-1
		}
	}
	return nums[l]
}