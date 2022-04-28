package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascendant := true
	descendant := true

	for i := 0; i < len(a)-2; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascendant = false
		} else if f(a[i], a[i+1]) < 0 {
			descendant = false
		}
	}

	if !descendant && !ascendant {
		return false
	} else {
		return true
	}
}
