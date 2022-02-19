package n0169

// time:O(n),space:O(1)
func majorityElement(nums []int) int {
    tmpMajority := nums[0]
    count := 1
    for i:=1;i<len(nums);i++{
        if nums[i]==tmpMajority {
            count++
        }else {
            count--
            if count==0 {
                tmpMajority=nums[i]
                count=1
            }
        }
    }
    return tmpMajority
}