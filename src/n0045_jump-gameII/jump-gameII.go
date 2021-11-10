package n0045

// 将整个数组划分为多个窗口，同一个窗口内的位置steps是一致的
// 当遍历到窗口右边界，更新窗口
func jump(nums []int) int {
	rightmost := 0 // 表示当前能够跳跃到的最右位置
	end := 0       // 记录边界位置
	steps := 0     // answer
	for i := 0; i < len(nums)-1; i++ {
		rightmost = max(rightmost, i+nums[i])
		if i == end {
			end = rightmost
			steps++
		}
	}

	return steps
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
