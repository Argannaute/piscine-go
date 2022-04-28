package piscine

func Map(f func(int) bool, a []int) []bool {
	var answer []bool
	for _, val := range a {
		answer = append(answer, f(val))
	}
	return answer
}
