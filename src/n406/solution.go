package n0406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.SliceStable(people, func(i, j int) bool {
		if people[i][0] != people[j][0] {
			return people[i][0] < people[j][0]
		} else {
			return people[i][1] > people[j][1]
		}
	})

	ans := make([][]int, len(people))
	for _, v := range people {
		count := 0
		ind := 0
		for {
			if ans[ind] == nil {
				if count==v[1] {
					ans[ind]=v
					break
				}else {
					count++
				}
			}
			ind++
		}
	}
	return ans
}
