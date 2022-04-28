package piscine

func Atoi(s string) int {
	result := 0
	sign := 1
	if s == "" || s == "0" {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + (int(ch) - '0')
	}
	return (result * sign)
}
