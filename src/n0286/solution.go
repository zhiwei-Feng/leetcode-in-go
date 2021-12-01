package n0286

func wallsAndGates(rooms [][]int) {
	dxdy := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]bool, len(rooms))
	for k := 0; k < len(rooms); k++ {
		visited[k] = make([]bool, len(rooms[0]))
	}
	queue := [][2]int{}
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == 0 {
				visited[i][j]=true
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	dis := 0
	for len(queue) > 0 {
		posNum := len(queue)
		for p := 0; p < posNum; p++ {
			top := queue[0]
			queue = queue[1:]
			rooms[top[0]][top[1]] = dis
			// visited[top[0]][top[1]] = true
			for _, d := range dxdy {
				newX, newY := top[0]+d[0], top[1]+d[1]
				if newX >= 0 && newX < len(rooms) && newY >= 0 && newY < len(rooms[0]) {
					if !visited[newX][newY] && rooms[newX][newY] != 0 && rooms[newX][newY] != -1 {
						visited[newX][newY] = true
						queue = append(queue, [2]int{newX, newY})
					}
				}
			}
		}
		dis++
	}
}
