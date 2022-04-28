package piscine

func IsLower(s string) bool {
	for _, letters := range s {
		if letters >= 97 && letters <= 122 {
		} else {
			return false
		}
	}
	return true
}
