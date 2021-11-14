package n0300longestincreasingsubseq

import "fmt"

func lengthOfLIS(nums []int) int {
	d := make([]int, len(nums)+1)
	longest := 1
	d[longest] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > d[longest] {
			longest++
			d[longest] = nums[i]
		} else {
			// 查找最后一个d[lo]<nums[i]的下一个位置
			// 换句话说就是找d[lo]>=nums[i]的下届
			lo, hi := 1, longest
			// pos := 0
			for lo < hi {
				m := lo + (hi-lo)/2
				if d[m] < nums[i] {
					lo=m+1
				} else {
					// pos = m
					hi = m
				}
			}
			d[lo] = min(d[lo], nums[i])
		}
	}
	return longest
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// dp版本，且可以返回对应的最长递增子序列
func lengthOfLIS_n2(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	ans_ind := 0
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	prev := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prev[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					prev[i] = j
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > ans {
			ans_ind = i
			ans = dp[i]
		}
	}
	res_list := []int{}
	for ans_ind != -1 {
		res_list = append([]int{nums[ans_ind]}, res_list...)
		ans_ind = prev[ans_ind]
	}
	fmt.Println(res_list)
	return ans
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
