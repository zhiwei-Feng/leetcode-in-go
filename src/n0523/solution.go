package n0523

func checkSubarraySum(nums []int, k int) bool {
	if len(nums)<2 {
		return false
	}
	remainMap := make(map[int]int)
	remainMap[0]=-1
	rem := 0
	for i:=0;i<len(nums);i++{
		rem = (rem + nums[i])%k
		if idx, ok := remainMap[rem];ok {
			if i-idx>=2 {
				return true
			}
		}else{
			remainMap[rem] = i
		}
	}
	return false
}