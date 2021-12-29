package n0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	inCount := make(map[int]int)
	connection := make(map[int][]int)
	q := []int{}
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
		for _, v := range connection[qt] {
			inCount[v]--
			if inCount[v] == 0 {
				delete(inCount, v)
				q = append(q, v)
			}
		}
	}

	return len(inCount) == 0
}
