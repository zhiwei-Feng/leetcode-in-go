package offer0057_2

// 滑动窗口解
func findContinuousSequence(target int) [][]int {
	winStart := 1
	winSum := 0
	ans := [][]int{}
	for winEnd := 1; winEnd <= target/2+1; winEnd++ {
		winSum += winEnd
		for winSum > target {
			winSum -= winStart
			winStart++
		}
		if winSum == target {
			tmp := []int{}
			for i := winStart; i <= winEnd; i++ {
				tmp = append(tmp, i)
			}
			ans = append(ans, tmp)
		}
	}
	return ans
}
