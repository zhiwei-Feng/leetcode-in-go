package n1345

func minJumps(arr []int) int {
	n := len(arr)
	idx := map[int][]int{} // 每个值的所有位置
	for i, v := range arr {
		idx[v] = append(idx[v], i)
	}

	vis := map[int]bool{0: true}
	type pair struct{ idx, step int }
	q := []pair{{0, 0}}
	for {
		qt := q[0]
		q = q[1:]
		i, step := qt.idx, qt.step
		if i == n-1 {
			return step
		}
		for _, j := range idx[arr[i]] {
			if !vis[j] {
				vis[j] = true
				q = append(q, pair{j, step + 1})
			}
		}
		// 注意，因为对于一个值x的所有idx，我们遍历到第一个下标位置的时候，下一次入队会将所有的idx入队
		// 因此，为了避免遍历到第二个下标位置进行重复判断的无用计算，这里删除掉这一组idx
		delete(idx, arr[i])
		if !vis[i+1] {
			// 不用判断i+1<n的原因是当i+1>=n时会在前面跳出循环
			vis[i+1] = true
			q = append(q, pair{i + 1, step + 1})
		}
		if i > 0 && !vis[i-1] {
			q = append(q, pair{i - 1, step + 1})
		}
	}
}

// func minJumps(arr []int) int {
// 	valueGroup := make(map[int][]int)
// 	for i, v := range arr {
// 		valueGroup[v] = append(valueGroup[v], i)
// 	}
// 	graph := make([][]int, len(arr))
// 	for i := range graph {
// 		graph[i] = make([]int, len(arr))
// 	}
// 	for i, v := range arr {
// 		if i-1 >= 0 {
// 			graph[i][i-1] = 1
// 			graph[i-1][i] = 1
// 		}
// 		if i+1 < len(arr) {
// 			graph[i][i+1] = 1
// 			graph[i+1][i] = 1
// 		}
// 		for _, other := range valueGroup[v] {
// 			if other != i {
// 				graph[i][other] = 1
// 				graph[other][i] = 1
// 			}
// 		}
// 	}

// 	q := []int{}
// 	q = append(q, 0)
// 	visit := make(map[int]bool)
// 	count := 0
// 	for len(q) != 0 {
// 		qt := q[0]
// 		q = q[1:]
// 		visit[qt] = true
// 		if qt == len(arr)-1 {
// 			return count
// 		}
// 		for i := 0; i < len(arr); i++ {
// 			if i == qt {
// 				continue
// 			}
// 			if visit[i] {
// 				continue
// 			}
// 			if graph[qt][i] == 0 {
// 				continue
// 			}
// 			q = append(q, i)
// 		}
// 		count++
// 	}
// 	return -1
// }
