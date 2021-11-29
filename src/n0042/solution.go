package n0042

// two pointer
func trap(height []int) int {
	leftMax := 0
	rightMax := 0
	left, right := 0, len(height)-1
	ans := 0
	for left <= right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			// 对于每一个right，找到一个left位置的height比rightMax大
			// 此时最大雨水量由左侧的leftMax决定
			// 对于每一个left，同理
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}

func trap_dp(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for j := len(height) - 2; j >= 0; j-- {
		rightMax[j] = max(rightMax[j+1], height[j])
	}

	ans := 0
	for i := 0; i < len(height); i++ {
		if water := min(leftMax[i], rightMax[i]) - height[i]; water > 0 {
			ans += water
		}
	}

	return ans
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
func min(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}
	return i2
}
