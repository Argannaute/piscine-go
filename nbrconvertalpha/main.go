package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var letter rune
	var cap rune
	result := 0
	for _, w := range os.Args[1:] {
		if os.Args[1] == "--upper" {
			cap += 32
		}
		//	if w[i] < 48 && w[i] > 57 {
		//		z01.PrintRune(' ')
		//	}
		if len(w) != 1 {
			for _, ch := range w {
				result = result*10 + (int(ch) - '0')
			}
		}
		for range w {
			letter += rune(result + 48)
			z01.PrintRune(letter)
			letter = 0
		}

	}
}
