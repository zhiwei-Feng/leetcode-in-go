package n0315

func countSmaller(nums []int) []int {
	var (
		n         = len(nums)
		index     = make([]int, n)
		temp      = make([]int, n)
		tempIndex = make([]int, n)
		ans       = make([]int, n)
		mergeSort func(start, end int)
	)
	for i := 0; i < n; i++ {
		index[i] = i
	}

	mergeSort = func(start, end int) {
		if start >= end {
			return
		}
		mid := start + (end-start)/2
		mergeSort(start, mid)
		mergeSort(mid+1, end)

		i, j := start, mid+1
		p := start
		for i <= mid && j <= end {
			if nums[i] <= nums[j] {
				temp[p] = nums[i]
				tempIndex[p] = index[i]
				ans[index[i]] += j - (mid + 1)
				i++
				p++
			} else {
				temp[p] = nums[j]
				tempIndex[p] = index[j]
				j++
				p++
			}
		}

		for i <= mid {
			temp[p] = nums[i]
			tempIndex[p] = index[i]
			ans[index[i]] += j - (mid + 1)
			i++
			p++
		}

		for j <= end {
			temp[p] = nums[j]
			tempIndex[p] = index[j]
			j++
			p++
		}

		for k := start; k <= end; k++ {
			index[k] = tempIndex[k]
			nums[k] = temp[k]
		}
	}

	l, r := 0, n-1
	mergeSort(l, r)
	return ans
}
