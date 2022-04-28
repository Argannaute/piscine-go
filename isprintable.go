package piscine

func IsPrintable(s string) bool {
	for _, letters := range s {
		if letters >= 32 && letters <= 126 {
		} else {
			return false
		}
	}
	return true
}
