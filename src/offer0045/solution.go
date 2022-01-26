package offer0045

import (
	"sort"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	sort.SliceStable(nums, func(i, j int) bool {
		// 取出第一位
		topi, topj := 0, 0
		ti := nums[i]
		tj := nums[j]
		for ti > 0 {
			topi = ti
			ti /= 10
		}
		for tj > 0 {
			topj = tj
			tj /= 10
		}
		if topi != topj {
			return topi < topj
		}

		// 第一位相等，则比较拼接后的字符串
		s1 := strconv.Itoa(nums[i]) + strconv.Itoa(nums[j])
		s2 := strconv.Itoa(nums[j]) + strconv.Itoa(nums[i])
		return s1 < s2
	})

	var ans strings.Builder
	for _, v := range nums {
		ans.WriteString(strconv.Itoa(v))
	}
	return ans.String()
}
