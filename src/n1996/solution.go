package n1996

import "sort"

func numberOfWeakCharacters(properties [][]int) (ans int) {
	sort.Slice(properties, func(i, j int) bool {
		p, q := properties[i], properties[j]
		return p[0] < q[0] || p[0] == q[0] && p[1] > q[1]
	})
	st := []int{}
	for _, p := range properties {
		for len(st) > 0 && st[len(st)-1] < p[1] {
			// 因为排序时，相同的attack按照defense从大到小排列
			// 所以p[1]一定会小于相同attack的栈顶，因此不计数
			st = st[:len(st)-1]
			ans++
		}
		st = append(st, p[1])
	}
	return
}
