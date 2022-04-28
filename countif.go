package piscine

func CountIf(f func(string) bool, tab []string) int {
	var count int
	for _, val := range tab {
		if f(val) {
			count++
		}
	}
	return count
}
