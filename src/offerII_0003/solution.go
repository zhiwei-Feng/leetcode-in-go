package offerii0003

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	ans := make([]int, n+1)
	ans[0] = 0
	ans[1] = 1
	prev := 0
	for i := 2; i <= n; i++ {
		if i&(i-1) == 0 { // i&(i-1)剔除最右边的1, 如果该i是2的指数倍, 则仅有一个1
			ans[i] = 1
			prev = i
		} else {
			ans[i] = 1 + ans[i-prev]
		}
	}
	return ans
}
