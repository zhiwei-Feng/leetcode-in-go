package offer0051

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

// 返回归并排序后有多少个逆序对
func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	i, j := start, mid+1
	tmp := []int{}
	for i <= mid && j <= end {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1)
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	for i <= mid {
		tmp = append(tmp, nums[i])
		cnt += end + 1 - (mid + 1)
		i++
	}
	for j <= end {
		tmp = append(tmp, nums[j])
		j++
	}
	for k := start; k <= end; k++ {
		nums[k] = tmp[k-start]
	}
	return cnt
}
