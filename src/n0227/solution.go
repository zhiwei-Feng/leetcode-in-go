package n0227

import ()

func calculate(s string) int {
	priority := map[byte]int{
		'-': 1,
		'+': 1,
		'*': 2,
		'/': 2,
		// '%': 2,
		// '^': 3,
	}
	nums := []int{0}
	ops := []byte{}
	idx := 0
	for idx < len(s) {
		c := s[idx]
		if c == ' ' {
			idx++
			continue
		}
		if c == '(' {
			ops = append(ops, c)
		} else if c == ')' {
			for len(ops) != 0 {
				if ops[len(ops)-1] != '(' {
					calc(&nums, &ops)
				} else {
					ops = ops[:len(ops)-1]
					break
				}
			}
		} else {
			if c >= '0' && c <= '9' {
				u := 0
				j := idx
				for j < len(s) && s[j] >= '0' && s[j] <= '9' {
					u = u*10 + int(s[j]-'0')
					j++
				}
				nums = append(nums, u)
				idx = j - 1
			} else {
				if idx > 0 && (s[idx-1] == '(' || s[idx-1] == '+' || s[idx-1] == '-') {
					nums = append(nums, 0)
				}
				for len(ops) > 0 && ops[len(ops)-1] != '(' {
					prev := ops[len(ops)-1]
					if priority[prev] >= priority[c] {
						calc(&nums, &ops)
					} else {
						break
					}
				}
				ops = append(ops, c)
			}
		}
		idx++
	}
	for len(ops) > 0 {
		calc(&nums, &ops)
	}
	return nums[len(nums)-1]
}

func calc(nums *[]int, ops *[]byte) {
	if len(*nums) == 0 || len(*nums) < 2 {
		return
	}
	if len(*ops) == 0 {
		return
	}
	b, a := PopNum(nums), PopNum(nums)
	op := PopOps(ops)
	ans := 0
	switch op {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		ans = a / b
	}
	*nums = append(*nums, ans)
}

func PopNum(nums *[]int) int {
	x := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	return x
}
func PopOps(ops *[]byte) byte {
	x := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]
	return x
}
