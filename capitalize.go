package piscine

func Capitalize(s string) string {
	answer := ""
	caps := false
	for i, letters := range s {
		a := s[i]
		if letters >= 48 && letters <= 57 || letters >= 65 && letters <= 90 || letters >= 97 && letters <= 122 {
			if caps {
				if letters >= 97 && letters <= 122 {
					answer += (string(a - 32))
					caps = false
				} else {
					answer += string(letters)
					caps = false
				}
			} else if i == 0 {
				if letters >= 97 && letters <= 122 {
					answer += string(letters - 32)
				} else {
					answer += string(letters)
				}
			} else if letters >= 65 && letters <= 90 {
				answer += (string(a + 32))
			} else {
				answer += string(letters)
			}
		} else {
			answer += string(letters)
			caps = true
		}
	}
	return answer
}
