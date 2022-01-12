package n1629


func slowestKey(releaseTimes []int, keysPressed string) byte {
	durationMap := make(map[int]byte)
	maxDuration := 0
	for i := 0; i < len(releaseTimes); i++ {
		var (
			curKey      byte
			curDuration int
		)
		curKey = keysPressed[i]
		if i == 0 {
			curDuration = releaseTimes[0]
		} else {
			curDuration = releaseTimes[i] - releaseTimes[i-1]
		}
		if v, ok := durationMap[curDuration];!ok {
			durationMap[curDuration] = curKey
		}else if v<curKey{
			durationMap[curDuration] = curKey
		}
		if curDuration > maxDuration {
			maxDuration = curDuration
		}
	}

	return durationMap[maxDuration]
}
