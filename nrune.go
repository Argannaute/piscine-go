package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	sLen := len(s)
	sRune := []rune(s)
	if sLen < n {
		return 0
	}
	answer := sRune[n-1]
	return answer
}
