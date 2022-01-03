package n0995

func minKBitFlips(nums []int, k int) (ans int) {
	n := len(nums)
	diff := make([]int, n+1) // diff[i]表示i-1 和 i位置的翻转次数差，若i-1多翻转一次，diff[i]=-1
	revCnt := 0              // 当前需要翻转的次数
	for i, v := range nums {
		revCnt += diff[i]
		if (v+revCnt)%2 == 0 { // 1和奇数次翻转 == 0和偶数次翻转 == 0%2 == 0
			if i+k > n {
				return -1
			}
			ans++
			revCnt++
			diff[i+k]--
		}
	}
	return
}
