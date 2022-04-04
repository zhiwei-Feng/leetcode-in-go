package n0862

func shortestSubarray(nums []int, k int) int {
	N := len(nums)
	P := make([]int, N+1) // 表示前缀和, P[i]表示前i个数之和
	q := make([]int, 0)   // 存储P[x]<=P[y]-k中x的最大值
	ans := N + 1
	for i := 0; i < N; i++ {
		P[i+1] = P[i] + nums[i]
		if nums[i] >= k {
			return 1
		}
	}

	for i := 0; i < len(P); i++ {
		// 1. 当前的前缀和P[i]<=P[x](i>x)，则显然x不再是最大的了，应剔除掉
		for len(q) != 0 && P[i] <= P[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		// 2. 不断缩减q，使得当P[i]-P[0]<k
		for len(q) != 0 && P[i]-P[q[0]]>=k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		q = append(q, i)
	}
	if ans == N+1 {
		return -1
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
