package piscine

func IsUpper(s string) bool {
	a := 0
	for _, letters := range s {
		if letters >= 65 && letters <= 90 {
			a++
		} else {
			return false
		}
	}
	return true
}
