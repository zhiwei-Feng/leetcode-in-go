package offer0038

import "sort"

func reverse(a []byte) {
    for i, n := 0, len(a); i < n/2; i++ {
        a[i], a[n-1-i] = a[n-1-i], a[i]
    }
}

func nextPermutation(nums []byte) bool {
    n := len(nums)
    i := n - 2
    for i >= 0 && nums[i] >= nums[i+1] {
        i--
    }
    if i < 0 {
        return false
    }
    j := n - 1
    for j >= 0 && nums[i] >= nums[j] {
        j--
    }
    nums[i], nums[j] = nums[j], nums[i]
    reverse(nums[i+1:])
    return true
}

func permutation(s string) (ans []string) {
    t := []byte(s)
    sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
    for {
        ans = append(ans, string(t))
        if !nextPermutation(t) {
            break
        }
    }
    return
}

func permutation_backtrace(s string) []string {
	ansMap := make(map[string]bool)
	ans := []string{}
	n := len(s)
	perm := make([]byte, n)
	vis := make([]bool, n)
	var backtrace func(int)
	backtrace = func(i int) {
		if i == n {
			ansMap[string(perm)] = true
			return
		}

		for j, b := range vis {
			if b {
				continue
			}
			vis[j] = true
			perm = append(perm, s[j])
			backtrace(i + 1)
			perm = perm[:len(perm)-1]
			vis[j] = false
		}
	}

	backtrace(0)
	for k, _ := range ansMap {
		ans = append(ans, k)
	}
	return ans
}
