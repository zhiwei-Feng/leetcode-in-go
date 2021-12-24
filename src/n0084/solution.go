package n0084

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func (s Stack) Top() int {
	return s[len(s)-1]
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}
	ans := 0

	var stack Stack
	for i := 0; i < n; i++ {
		for len(stack) != 0 && heights[stack.Top()] > heights[i] {
			curHeight := heights[stack.Pop()]
			for len(stack) != 0 && curHeight == heights[stack.Top()] {
				stack.Pop()
			}

			var curWidth int
			if len(stack) == 0 {
				curWidth = i
			} else {
				curWidth = i - stack.Top() - 1
			}

			area := curWidth * curHeight
			ans = max(ans, area)
		}

		stack.Push(i)
	}

	for len(stack) != 0 {
		curHeight := heights[stack.Pop()]
		for len(stack) != 0 && curHeight == heights[stack.Top()] {
			stack.Pop()
		}

		var curWidth int
		if len(stack) == 0 {
			curWidth = n
		} else {
			curWidth = n - stack.Top() - 1
		}

		area := curWidth * curHeight
		ans = max(ans, area)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
