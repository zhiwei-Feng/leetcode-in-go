package offer0031

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) == 0 {
		return true
	}
	stack := []int{}
	popi := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[popi] {
			stack = stack[:len(stack)-1]
			popi++
		}
	}
	return len(stack) == 0
}
