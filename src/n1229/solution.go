package n1229

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	timeLine := make([][2]int, 0, len(slots1)+len(slots2))
	for _, x := range slots1 {
		timeLine = append(timeLine, [2]int{x[0], 0})
		timeLine = append(timeLine, [2]int{x[1], 2})
	}
	for _, y := range slots2 {
		timeLine = append(timeLine, [2]int{y[0], 1})
		timeLine = append(timeLine, [2]int{y[1], 3})
	}

	sort.Slice(timeLine, func(i, j int) bool {
		return timeLine[i][0] < timeLine[j][0]
	})
	var s1, s2, cnt int
	ans := []int{}
	for _, x := range timeLine {
		if x[1] == 0 {
			cnt++
			s1 = x[0]
		}
		if x[1] == 1 {
			cnt++
			s2 = x[0]
		}
		if x[1] == 2 || x[1] == 3 {
			if cnt == 2 && max(s1, s2)+duration <= x[0] {
				ans = append(ans, max(s1, s2))
				ans = append(ans, max(s1, s2)+duration)
				break
			}
			cnt--
		}
	}
	return ans
}

func minAvailableDuration_twopointer(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.SliceStable(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.SliceStable(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})
	var (
		i, j = 0, 0
	)

	for i < len(slots1) && j < len(slots2) {
		left := max(slots1[i][0], slots2[j][0])
		right := min(slots1[i][1], slots2[j][1])

		if left <= right && right-left >= duration {
			return []int{left, left + duration}
		} else {
			if slots1[i][1] < slots2[j][1] {
				i++
			} else {
				j++
			}
		}
	}

	return []int{}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
