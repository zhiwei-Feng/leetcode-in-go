package offer0046

import "strconv"

func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1 // p表示两位作为当前数的累积，q表示一位作为当前的累积，r表示总共的
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
