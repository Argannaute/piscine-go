package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	s := string(n)
	for _, val := range s {
		if val == '-' {
			z01.PrintRune('-')
		} else {
			z01.PrintRune(val)
		}
	}
}
