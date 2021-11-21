package n0202

import "math"

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	slow, fast := n, n
	for fast != 1 {
		slow = happyChange(slow)
		fast = happyChange(happyChange(fast))
		if slow == fast && fast != 1 {
			return false
		}
	}
	return true
}

func isHappy_hashmap(n int) bool {
	if n == 1 {
		return true
	}
	history := make(map[int]int)
	history[n]++
	for n != 1 || n > math.MaxInt32 {
		n = happyChange(n)
		if _, ok := history[n]; ok {
			return false
		}
		history[n]++
	}
	return true
}

func happyChange(n int) int {
	ans := 0
	for n > 0 {
		tmp := n % 10
		ans += tmp * tmp
		n /= 10
	}
	return ans
}
