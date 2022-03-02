package n0053

func maxSubArray(nums []int) int {
	return get(nums, 0, len(nums)-1).Msum
}

func pushUp(l, r Status) Status {
	merge := new(Status)
	merge.Lsum = max(l.Lsum, l.Sum+r.Lsum)
	merge.Rsum = max(r.Rsum, r.Sum+l.Rsum)
	merge.Msum = max(max(l.Msum, r.Msum), l.Rsum+r.Lsum)
	merge.Sum = l.Sum+r.Sum
	return *merge
}

func get(nums []int, l, r int) Status {
	if l==r {
		// 一个元素
		return Status{nums[l], nums[l],nums[l],nums[l]}
	}

	m := l + (r-l)/2
	lStatus := get(nums, l, m)
	rStatus := get(nums, m+1, r)
	return pushUp(lStatus, rStatus)
}

type Status struct {
	Lsum int // 左端点开始的最大子段和
	Rsum int // 右端点开始的最大子段和
	Msum int // 区间内最大子段和
	Sum int // 区间和
}

// func maxSubArray(nums []int) int {
// 	ans := 0
// 	tmpMax := 0
// 	for i, v := range nums {
// 		if i == 0 {
// 			tmpMax = v
// 		} else {
// 			tmpMax = max(nums[i], tmpMax+nums[i])
// 		}
// 		if tmpMax > ans {
// 			ans = tmpMax
// 		}
// 	}
// 	return ans
// }

// func maxSubArray(nums []int) int {
// 	dp := make([]int, len(nums))
// 	dp[0] = nums[0]
// 	ans := dp[0]
// 	for i := 1; i < len(nums); i++ {
// 		dp[i] = max(nums[i], dp[i-1]+nums[i])
// 		if dp[i] > ans {
// 			ans = dp[i]
// 		}
// 	}

// 	return ans
// }

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
