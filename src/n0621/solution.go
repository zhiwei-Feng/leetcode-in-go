package n0621

func leastInterval(tasks []byte, n int) (minTime int) {
	minLen := len(tasks)
	cnt := make(map[byte]int)
	for _, t := range tasks {
		cnt[t]++
	}
	maxNum := 0
	for _, v := range cnt {
		if v > maxNum {
			maxNum = v
		}
	}
	sameNumTasks := 0
	for _, v := range cnt {
		if v == maxNum {
			sameNumTasks++
		}
	}
	time := (maxNum-1)*(n+1) + sameNumTasks
	if time<minLen{
		return minLen
	}
	return time
}
