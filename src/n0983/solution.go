package n0983

import "math"


func mincostTickets(days []int, costs []int) int {
	n := len(days)
	memo := make([]int, n)
	for i:=0;i<n;i++{
		memo[i] = math.MaxInt32
	}
	duration := []int{1, 7, 30}
	var dp func(pos int) int
	dp = func(i int) int {
		if i >= n {
			return 0
		}
		if memo[i] != math.MaxInt32 {
			return memo[i]
		}
		j := i
		for k := 0; k < 3; k++ {
			for ; j < n && days[j] < days[i]+duration[k]; j++ {
			}

			memo[i] = min(memo[i], dp(j)+costs[k])
		}

		return memo[i]
	}

	return dp(0)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
