package piscine

func IsNumeric(s string) bool {
	for _, letters := range s {
		if letters >= 48 && letters <= 57 {
		} else {
			return false
		}
	}
	return true
}
