package offerii0002

func addBinary(a string, b string) string {
	ans := make([]byte, 0, len(a)+len(b))
	i, j := len(a)-1, len(b)-1
	carry := 0
	for i >= 0 && j >= 0 {
		sum := carry + int(a[i]-'0') + int(b[j]-'0')
		carry = sum / 2
		ans = append(ans, byte('0'+sum%2))
		i--
		j--
	}
	for i >= 0 {
		sum := carry + int(a[i]-'0')
		carry = sum / 2
		ans = append(ans, byte('0'+sum%2))
		i--
	}
	for j >= 0 {
		sum := carry + int(b[j]-'0')
		carry = sum / 2
		ans = append(ans, byte('0'+sum%2))
		j--
	}
	if carry > 0 {
		ans = append(ans, byte('0'+carry))
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return string(ans)
}
