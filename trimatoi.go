package piscine

func TrimAtoi(s string) int {
	a := 0
	neg := false
	for _, numbers := range s {
		if numbers >= 48 && numbers <= 57 {
			a = (a * 10) + (int(numbers) - 48)
		}
		if a == 0 {
			if numbers == 45 {
				neg = true
			}
		}
	}
	if neg {
		return (a * -1)
	}
	return a
}
