package n0070

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	prev1, prev2 := 1, 1
	for i := 2; i <= n; i++ {
		prev1, prev2 = prev2, prev1+prev2
	}
	return prev2
}
