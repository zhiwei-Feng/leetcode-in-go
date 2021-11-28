package n1182

import "sort"

// left is the minimum distance of colors in the left side
// the same as right
func shortestDistanceColor(colors []int, queries [][]int) []int {
	ans := []int{}
	left, right := make([][3]int, len(colors)), make([][3]int, len(colors))
	for i := 0; i < len(colors); i++ {
		left[i] = [3]int{-1, -1, -1}
		right[i] = [3]int{-1, -1, -1}
	}

	//get left
	for i := 0; i < len(colors); i++ {
		for j := 0; j < 3; j++ {
			if i-1 >= 0 && left[i-1][j] != -1 {
				left[i][j] = left[i-1][j] + 1
			}
		}
		left[i][colors[i]-1] = 0
	}
	// get right
	for i := len(colors) - 1; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			if i+1 < len(colors) && right[i+1][j] != -1 {
				right[i][j] = right[i+1][j] + 1
			}
		}
		right[i][colors[i]-1] = 0
	}

	for _, q := range queries {
		leftDis := left[q[0]][q[1]-1]
		rightDis := right[q[0]][q[1]-1]
		tmp := 2<<16 - 1
		if leftDis != -1 {
			tmp = min(tmp, leftDis)
		}
		if rightDis != -1 {
			tmp = min(tmp, rightDis)
		}

		if tmp != 2<<16-1 {
			ans = append(ans, tmp)
		} else {
			ans = append(ans, -1)
		}
	}

	return ans
}

func shortestDistanceColor_binary_search(colors []int, queries [][]int) []int {
	colorIndexMap := make(map[int][]int)
	ans := []int{}
	for i, color := range colors {
		colorIndexMap[color] = append(colorIndexMap[color], i)
	}

	for _, query := range queries {
		indList := colorIndexMap[query[1]]
		if len(indList) == 0 {
			ans = append(ans, -1)
			continue
		}

		possible_idx := sort.SearchInts(indList, query[0])
		possible_dis := 2<<16 - 1
		if possible_idx < len(indList) {
			possible_dis = min(possible_dis, indList[possible_idx]-query[0])
		}
		if possible_idx-1 >= 0 {
			possible_dis = min(possible_dis, query[0]-indList[possible_idx-1])
		}
		ans = append(ans, possible_dis)
	}
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
