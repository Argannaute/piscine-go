package piscine

func LastRune(s string) rune {
	sRune := []rune(s)
	lsRune := len(sRune)
	a := lsRune - 1
	return sRune[a]
}
