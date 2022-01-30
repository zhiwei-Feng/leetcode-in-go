package n1342

func numberOfSteps(num int) int {
	return recursive(num, make(map[int]int))
}

func recursive(num int, mem map[int]int) int {
	if num == 0 {
		return 0
	}
	if v, ok := mem[num]; ok {
		return v
	}
	if num%2 == 0 {
		ans := recursive(num/2, mem) + 1
		mem[num] = ans
		return ans
	} else {
		ans := recursive(num-1, mem) + 1
		mem[num] = ans
		return ans
	}
}
