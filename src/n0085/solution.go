package n0085

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	left := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				left[j] += 1
			} else {
				left[j] = 0
			}
		}

		ans = max(ans, largestRectangleArea(left))
	}

	return ans
}

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

// func maximalRectangle_m2n(matrix [][]byte) (ans int) {
// 	if len(matrix) == 0 {
// 		return
// 	}
// 	m, n := len(matrix), len(matrix[0])
// 	left := make([][]int, m)
// 	for i, row := range matrix {
// 		left[i] = make([]int, n)
// 		for j, v := range row {
// 			if v == '0' {
// 				continue
// 			}
// 			if j == 0 {
// 				left[i][j] = 1
// 			} else {
// 				left[i][j] = left[i][j-1] + 1
// 			}
// 		}
// 	}
// 	for i, row := range matrix {
// 		for j, v := range row {
// 			if v == '0' {
// 				continue
// 			}
// 			width := left[i][j]
// 			area := width
// 			for k := i - 1; k >= 0; k-- {
// 				width = min(width, left[k][j])
// 				area = max(area, (i-k+1)*width)
// 			}
// 			ans = max(ans, area)
// 		}
// 	}
// 	return
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
