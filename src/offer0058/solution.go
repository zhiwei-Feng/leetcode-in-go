package offer0058

func reverseLeftWords(s string, n int) string {
	toByte := []byte{}
	for i := 0; i < len(s); i++ {
		toByte = append(toByte, s[i])
	}

	reverse(toByte[:n])
	reverse(toByte[n:])
	reverse(toByte)
	return string(toByte)
}

func reverse(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
