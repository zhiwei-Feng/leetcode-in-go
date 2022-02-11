package n0001

func twoSum(nums []int, target int) []int {
	exist := make(map[int]int) // val->index
	for i, v := range nums {
		if otheri, ok := exist[target-v]; ok {
			return []int{otheri, i}
		}
		exist[v] = i
	}
	return []int{}
}
