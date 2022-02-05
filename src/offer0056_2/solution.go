package offer0056_2

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, v := range nums {
		ones = ones ^ v&(-twos-1)
		twos = twos ^ v&(-ones-1)
	}

	return ones
}

// func singleNumber(nums []int) int {
// 	bitCount := [32]int{}
// 	for i := 0; i < len(nums); i++ {
// 		bitMask := 1
// 		for j := 31; j >= 0; j-- {
// 			if (bitMask & nums[i]) == bitMask {
// 				bitCount[j]++
// 			}
// 			bitMask <<= 1
// 		}
// 	}

// 	res := 0
// 	for i := 0; i < 32; i++ {
// 		res <<= 1
// 		res += bitCount[i] % 3
// 	}
// 	return res
// }
