package n0128

func longestConsecutive(nums []int) int {
	fa := make(map[int]int)
	for _, num := range nums {
		fa[num] = num
	}
	ans := 0
	for _, num := range nums {
		if _, ok := fa[num+1]; ok {
			union(num, num+1, fa)
		}
	}
	for _, num := range nums {
		right := find(num, fa)
		if right-num+1 > ans {
			ans = right - num + 1
		}
	}
	return ans
}

func find(x int, fa map[int]int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = find(fa[x], fa)
	return fa[x]
}

func union(i, j int, fa map[int]int) {
	x := find(i, fa)
	y := find(j, fa)
	fa[x] = y
}

// map存储当前数所在连续区间的最长的长度
func longestConsecutive_dp(nums []int) int {
	continueRangeMap := make(map[int]int)
	ans := 0
	for _, num := range nums {
		if _, ok := continueRangeMap[num]; !ok {
			left := continueRangeMap[num-1]
			right := continueRangeMap[num+1]
			curLen := left + right + 1
			if curLen > ans {
				ans = curLen
			}

			continueRangeMap[num] = curLen
			continueRangeMap[num-left] = curLen
			continueRangeMap[num+right] = curLen
		}
	}
	return ans
}

// 使用hashtable存储每个数字的连续右边界
func longestConsecutive_hash2(nums []int) int {
	rightMap := make(map[int]int)
	for _, num := range nums {
		rightMap[num] = num
	}
	ans := 0
	for _, num := range nums {

		if right, exist := rightMap[num]; exist {
			for {
				v, ok := rightMap[right+1]
				if ok {
					delete(rightMap, right)
					right = v
				} else {
					break
				}
			}
			rightMap[num] = right
			if right-num+1 > ans {
				ans = right - num + 1
			}
		}
	}
	return ans
}

// 使用hashtable存储出现的数字
func longestConsecutive_hash1(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}
	longestStreak := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}
			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}
	return longestStreak
}
