package piscine

func Unmatch(a []int) int {
	counter := 0
	failCount := 0
	for i := 0; i < len(a); i++ {
		for _, val := range a {
			if a[i] != val {
				counter++
			} else {
				failCount++
			}
		}
		if counter == len(a)-1 {
			return a[i]
		}
		if failCount%2 == 1 {
			return a[i]
		}
		failCount = 0
		counter = 0
	}
	return -1
}
