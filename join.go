package piscine

func Join(strs []string, sep string) string {
	a := ""
	answer := ""
	first := false
	for _, text := range strs {
		a = text
		if first {
			answer += sep
		}
		answer += a
		first = true
	}
	return answer
}
