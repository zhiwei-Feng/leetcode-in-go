package n0051nqueen

// 将棋盘从上往下看，行号记为pos，每一行的每一列不可以有重复占用
func solveNQueens(n int) [][]string {
	visit := make([]bool, n)   // 记录每一列是否有一行占用
	record := make([][]int, 0) // 记录已遍历的行和对应的占用列
	tmpBoard := []string{}     // 临时记录当前棋盘中皇后分布的情况
	ans := [][]string{}        // 最终所有的解决方案
	var dfs func(int)

	dfs = func(pos int) {
		if pos == n {
			// 完成全部遍历，将当前的棋盘情况添加进ans
			copyBoard := make([]string, len(tmpBoard))
			copy(copyBoard, tmpBoard)
			ans = append(ans, copyBoard)
			return
		}

		tmpLine := []byte{} // 记录每一行的字节，并全部初始化为'.'
		for i := 0; i < n; i++ {
			tmpLine = append(tmpLine, '.')
		}
		// 遍历该行的每一列
		// 1. 已被占用则跳出
		// 2. 如果和之前行的皇后处于攻击范围内则跳出
		for i := 0; i < n; i++ {
			if visit[i] {
				// 当前
				continue
			}
			conflict := false
			for j := 0; j < len(record); j++ {
				if abs(record[j][0]-pos) == abs(record[j][1]-i) {
					conflict = true
					break
				}
			}
			if conflict {
				continue
			}
			// set i
			tmpLine[i] = 'Q'
			visit[i] = true
			record = append(record, []int{pos, i})
			tmpBoard = append(tmpBoard, string(tmpLine))
			dfs(pos + 1)
			// 回溯处理
			tmpLine[i] = '.'
			visit[i] = false
			record = record[:len(record)-1]
			tmpBoard = tmpBoard[:len(tmpBoard)-1]
		}
	}
	dfs(0)

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
