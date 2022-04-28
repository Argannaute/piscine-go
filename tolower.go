package piscine

func ToLower(s string) string {
	answer := ""
	for _, letters := range s {
		if letters >= 65 && letters <= 90 {
			answer += string(letters + 32)
		}
		if letters >= 97 && letters <= 122 {
			answer += string(letters)
		}
		if letters < 65 || letters > 122 || letters > 90 && letters < 97 {
			answer += string(letters)
		}
	}
	return answer
}
