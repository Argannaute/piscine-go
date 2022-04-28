package piscine

func Max(a []int) int {
	ans := 0
	if len(a) == 0 {
		return 0
	}
	for i := 0; i < len(a); i++ {
		if ans < a[i] {
			ans = a[i]
		}
	}
	return ans
}
