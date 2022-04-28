package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	a := []rune(s)
	for _, s := range a {
		z01.PrintRune(s)
	}
}
