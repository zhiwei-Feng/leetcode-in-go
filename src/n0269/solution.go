package n0269

import "strings"

type void struct{}

func alienOrder(words []string) string {
	graph := make(map[byte]map[byte]void)
	n := len(words)

	// 1.遍历相邻字符串进行图的构建
	for i := 0; i < n-1; i++ {
		hasDiff := false
		for j := 0; j < min(len(words[i]), len(words[i+1])); j++ {
			if words[i][j] == words[i+1][j] {
				continue
			}
			hasDiff = true
			if _, ok := graph[words[i][j]]; !ok {
				graph[words[i][j]] = make(map[byte]void)
			}
			graph[words[i][j]][words[i+1][j]] = void{}
			break
		}

		// 处理 "abc","ab"的情况
		if !hasDiff && len(words[i]) > len(words[i+1]) {
			return ""
		}
	}

	// 初始化所有出现过的字符
	inDegree := make(map[byte]int)
	for _, str := range words {
		for i := range str {
			inDegree[str[i]] = 0
		}
	}
	// 2. 构建入度map
	for _, v := range graph {
		for tail := range v {
			inDegree[tail]++
		}
	}

	var builder strings.Builder
	queue := []byte{}
	for k, v := range inDegree {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	// 3. 拓扑排序
	for len(queue) != 0 {
		level := len(queue)
		for i := 0; i < level; i++ {
			cur := queue[0]
			queue = queue[1:]
			builder.WriteByte(cur)
			for k := range graph[cur] {
				inDegree[k]--
				if inDegree[k] == 0 {
					queue = append(queue, k)
				}
			}
		}
	}

	allDone := true
	for _, v := range inDegree {
		if v != 0 {
			allDone = false
			break
		}
	}

	if allDone {
		return builder.String()
	} else {
		return ""
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
