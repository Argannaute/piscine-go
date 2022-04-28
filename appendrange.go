package piscine

func AppendRange(min, max int) []int {
	var answer []int
	if min >= max {
		return []int(nil)
	}
	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	return answer
}
