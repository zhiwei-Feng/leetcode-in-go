package offer0010_2

func numWays(n int) int {
	const MO = 1e9 + 7
	if n < 2 {
		return 1
	}
	prev1, prev2 := 1, 1
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (prev1 + prev2) % MO
		prev1, prev2 = prev2, ans
	}
	return ans
}
