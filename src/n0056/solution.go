package n0056

import "sort"

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})

	ans := [][]int{}
	lastInd := -1
	for i := range intervals {
		if i == 0 {
			ans = append(ans, intervals[i])
			lastInd = 0
			continue
		}

		if intervals[i][0] <= ans[lastInd][1] {
			// overlap
			ans[lastInd][0] = min(ans[lastInd][0], intervals[i][0])
			ans[lastInd][1] = max(ans[lastInd][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
			lastInd++
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
