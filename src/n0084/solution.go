package n0084

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	ans := 0
	stack := []int{0}
	heights_extend := append([]int{0}, append(heights, 0)...)
	for i, v := range heights_extend {
		for len(stack) > 0 && heights_extend[stack[len(stack)-1]] > v {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := heights_extend[j] * (i - stack[len(stack)-1] - 1)
			if tmp > ans {
				ans = tmp
			}
		}
		stack = append(stack, i)
	}
	return ans
}
