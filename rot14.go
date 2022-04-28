package piscine

func Rot14(s string) string {
	var newChar rune
	newStr := ""
	for _, val := range s {
		newChar = val
		if val >= 97 && val <= 122 {
			if (val + 14) > 122 {
				newChar = (val - 26) + 14
			} else {
				newChar = val + 14
			}
		}
		if val >= 65 && val <= 90 {
			newChar = val + 14
			if newChar > 90 {
				newChar = newChar - 26
			}
		}
		newStr += string(newChar)
	}
	return newStr
}
