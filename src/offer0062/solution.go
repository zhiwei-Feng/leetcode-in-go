package offer0062

// 迭代版本
func lastRemaining(n int, m int) int {
	fn := 0
	for i := 2; i <= n; i++ {
		fn = (m + fn) % i
	}
	return fn
}

func lastRemaining_recur(n int, m int) int {
	return f(n, m)
}

// 对于n个数，删除后第m个数得到，返回最后一个数的坐标
func f(n, m int) int {
	if n == 1 {
		return 0
	}

	x := f(n-1, m) // 移动x+1次
	return (m - 1 + x + 1) % n
}
