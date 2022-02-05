package offer0066

// 第二次遍历可以精简空间为1
func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	ans := make([]int, len(a))
	ans[0] = 1
	for i := 1; i < len(a); i++ {
		ans[i] = ans[i-1] * a[i-1]
	}
	tmp := 1
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		ans[i] *= tmp
	}
	return ans
}
