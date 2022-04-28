package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	answer := make([]int, (max - min))
	for i, place := min, 0; i < max; i, place = i+1, place+1 {
		answer[place] = i
	}
	return answer
}
