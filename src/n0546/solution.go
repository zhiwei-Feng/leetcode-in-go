package n0546

func removeBoxes(boxes []int) int {
	n := len(boxes)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	return maxPoint(boxes, 0, n-1, mem)
}

func maxPoint(boxes []int, left, right int, mem [][]int) int {
	if left > right {
		return 0
	}
	if left == right {
		return 1
	}
	if mem[left][right] != 0 {
		return mem[left][right]
	}
	i := left
	ans := 0
	for i <= right {
		tmpAns := 0
		end := i
		for end<=right&&boxes[end] == boxes[i] {
			end++
		}
		// i..end-1
		tmpAns += maxPoint(boxes, left, i-1, mem)
		tmpAns += (end - i) * (end - i)
		tmpAns += maxPoint(boxes, end, right, mem)
		if tmpAns > ans {
			ans = tmpAns
		}
		i=end
	}

	mem[left][right] = ans
	return ans
}
