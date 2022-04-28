package piscine

func Index(s string, toFind string) int {
	Rtofind := []rune(toFind)
	length := len(toFind)
	slength := len(s)
	if slength == 0 || length == 0 {
		return 0
	}
	if length > slength {
		return 0
	}
	for index, letters := range s {
		if (index + length) > slength {
			return -1
		}
		if letters == Rtofind[0] {
			if s[index:(index+length)] == toFind {
				return index
			}
		}
	}
	return -1
}
