package n0210

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		inCount    = make(map[int]int)
		connection = make(map[int][]int)
		path       = []int{}
		q          = []int{}
	)

	for _, v := range prerequisites {
		inCount[v[0]]++
		connection[v[1]] = append(connection[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if _, ok := inCount[i]; !ok {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		qt := q[0]
		q = q[1:]
		path = append(path, qt)

		for _, v := range connection[qt] {
			inCount[v]--
			if inCount[v] == 0 {
				delete(inCount, v)
				q = append(q, v)
			}
		}
	}

	if len(inCount) > 0 {
		return []int{}
	} else {
		return path
	}
}
