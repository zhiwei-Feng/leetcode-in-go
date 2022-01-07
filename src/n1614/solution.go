package n1614

func maxDepth(s string) int {
	stack := []rune{}
	maxDepth:=0
	for _, v := range s {
		if v!='('&&v!=')'{
			continue
		}
		if v=='('{
			stack = append(stack, v)
			if len(stack)>maxDepth {
				maxDepth = len(stack)
			}
		}else {
			stack = stack[:len(stack)-1]
		}
	}

	return maxDepth
}