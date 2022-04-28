package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	firstrun := false
	for a := '0'; a < '8'; a++ {
		for b := '0'; b < '9'; b++ {
			for c := '0'; c <= '9'; c++ {
				if a < b && a != b && b < c {
					if firstrun {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						firstrun = true
					}
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
