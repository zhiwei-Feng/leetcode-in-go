package bytedance0002

import (
    "sort"
    "strconv"
)

// 给定一个数n如23121;给定一组数字a如[2 4 9]求由a中元素组成的小于n的最大数
// n=23121, nums=[2,4,9], ans = 22999
// n=21121, nums=[2,4,9], ans = 9999
func solve(n int, nums []int) int {
    n_str := strconv.Itoa(n)
    sort.Ints(nums)
    ans := -1
    // pass为true的情况，只有最高位没有数，或者上一位数小于对应n中的值
    var dfs func(idx int, pass bool, tmp int) bool 
    dfs = func(idx int, pass bool, tmp int) bool {
        if idx == len(n_str) {
            ans = tmp
            return true
        }
        if pass {
            return dfs(idx+1, pass, tmp*10+nums[len(nums)-1])
        }else {
            iNum := int(n_str[idx]-'0')
            for i:=len(nums)-1;i>=0;i--{
                if nums[i]==iNum {
                    if dfs(idx+1, false, tmp*10+nums[i]) {
                        return true
                    }
                }else if nums[i]<iNum {
                    if dfs(idx+1, true, tmp*10+nums[i]){
                        return true
                    }
                }else {
                    continue
                }
            }
            // 如果该位置都放不了且idx!=0 则return false
            if idx!=0 {
                return false
            }
            // idx == 0，则后续都是最大值
            return dfs(idx+1, true, tmp)
        }
    }

    dfs(0, false, 0)
    return ans
}