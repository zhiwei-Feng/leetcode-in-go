package offer0064

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(i int) bool {
		ans += i
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

// func sumNums(n int) int {
//     ans, A, B := 0, n, n + 1
//     addGreatZero := func() bool {
//         ans += A
//         return ans > 0
//     }

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1
//     return ans >> 1
// }
