package offer0050

type ByteCount struct {
	Pos  int
	Name byte
}

func firstUniqChar(s string) byte {
	n := len(s)
	byteCount := [26]int{}
	for i := 0; i < 26; i++ {
		byteCount[i] = n + 1
	}
	finalUniqByte := []ByteCount{}
	for _, v := range s {
		curPos := v - 'a'
		if byteCount[curPos] == n+1 {
			//first time
			finalUniqByte = append(finalUniqByte, ByteCount{int(curPos), byte(v)})
			byteCount[curPos] = 1
		} else {
			byteCount[curPos]++
			for len(finalUniqByte) > 0 && byteCount[finalUniqByte[0].Pos] != 1 {
				finalUniqByte = finalUniqByte[1:]
			}
		}
	}

	if len(finalUniqByte) > 0 {
		return finalUniqByte[0].Name
	}
	return ' '
}
