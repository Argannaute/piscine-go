package piscine

func IsAlpha(s string) bool {
	for _, letters := range s {
		if letters >= 97 && letters <= 122 {
		} else if letters >= 48 && letters <= 57 {
		} else if letters >= 65 && letters <= 90 {
		} else {
			return false
		}
	}
	return true
}
