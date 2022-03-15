package n0088

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	sortedIdx := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			nums1[sortedIdx] = nums2[j]
			sortedIdx--
			j--
		} else {
			nums1[sortedIdx] = nums1[i]
			sortedIdx--
			i--
		}
	}
	for i >= 0 {
		nums1[sortedIdx] = nums1[i]
		sortedIdx--
		i--
	}
	for j >= 0 {
		nums1[sortedIdx] = nums2[j]
		sortedIdx--
		j--
	}
}

// 去重版本
func merge_unique(nums1 []int, m int, nums2 []int, n int) []int {
	i, j := m-1, n-1
	sortedIdx := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] <= nums2[j] {
			if nums2[j] == nums1[sortedIdx] {
				j--
			} else {
				nums1[sortedIdx] = nums2[j]
				sortedIdx--
				j--
			}

		} else {
			if nums1[i] == nums1[sortedIdx] {
				i--
			} else {
				nums1[sortedIdx] = nums1[i]
				sortedIdx--
				i--
			}
		}
	}
	for i >= 0 {
		if nums1[sortedIdx] != nums1[i] {
			nums1[sortedIdx] = nums1[i]
			sortedIdx--
		}
		i--
	}
	for j >= 0 {
		if nums2[j] != nums1[sortedIdx] {
			nums1[sortedIdx] = nums2[j]
			sortedIdx--
		}
		j--
	}
	return nums1[sortedIdx+1:]
}
