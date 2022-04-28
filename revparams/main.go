package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[1:]
	nameLength := len(name) - 1
	for i := range name {
		for _, letters := range name[nameLength-i] {
			z01.PrintRune(letters)
		}
		z01.PrintRune('\n')
	}
}
