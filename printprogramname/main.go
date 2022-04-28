package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	for _, letters := range name[2:] {
		z01.PrintRune(letters)
	}
	z01.PrintRune('\n')
}
