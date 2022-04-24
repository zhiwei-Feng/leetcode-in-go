package offerii0006

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i, j}
		}else if numbers[i]+numbers[j] > target {
			j--
		}else {
			i++
		}
	}
	return []int{}
}
