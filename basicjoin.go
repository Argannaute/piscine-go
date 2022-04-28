package piscine

func BasicJoin(elems []string) string {
	a := ""
	answer := ""
	for _, text := range elems {
		a = text
		answer += a
	}
	return answer
}
