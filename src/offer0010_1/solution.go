package offer0010_1

func fib(n int) int {
	if n < 2 {
		return n
	}
    const MO = 1e9 + 7
	prev1, prev2 := 0, 1
	var ans int
	for i:=2; i <= n; i++ {
		ans = (prev1 + prev2) % MO
		prev1, prev2 = prev2, ans
	}
	return ans
}
