package n0539

import (
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	if len(timePoints) > 1440 {
		return 0
	}
	ans := 24 * 60
	sort.Strings(timePoints)
	prevTotalMinute := Duration(timePoints[0])
	for i := 1; i < len(timePoints); i++ {
		curTotalMinute := Duration(timePoints[i])
		ans = min(ans, curTotalMinute-prevTotalMinute)
		prevTotalMinute = curTotalMinute
	}

	special := 24*60 - (prevTotalMinute - Duration(timePoints[0]))
	if special < ans {
		ans = special
	}
	return ans
}

func Duration(p string) int {
	t := strings.Split(p, ":")
	HH, _ := strconv.Atoi(t[0])
	MM, _ := strconv.Atoi(t[1])
	return HH*60 + MM
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
