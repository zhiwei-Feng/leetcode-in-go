package n0020

func isValid(s string) bool {
	stack := []byte{}
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}

		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(stack) == 0 {
				return false
			}
			if match(stack[len(stack)-1], s[i]) {
				stack = stack[:len(stack)-1]
			}else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func match(b1, b2 byte) bool {
	return (b1 == '(' && b2 == ')') || (b1 == '[' && b2 == ']') || (b1 == '{' && b2 == '}')
}
