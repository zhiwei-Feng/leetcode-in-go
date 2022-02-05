package offer0039

// 该种解法不需要额外的空间
func majorityElement(nums []int) int {
	count := 0
	candidate := 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count++
		} else {
			if v != candidate {
				count--
			} else {
				count++
			}
		}
	}
	return candidate
}

// func majorityElement(nums []int) int {
// 	countMap := map[int]int{}
// 	for _, v := range nums {
// 		countMap[v]++
// 		if countMap[v] > len(nums)/2 {
// 			return v
// 		}
// 	}
// 	return -1
// }
