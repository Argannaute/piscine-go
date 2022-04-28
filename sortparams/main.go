package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for range os.Args[1:] {
		for i, v := range os.Args {
			if v == os.Args[len(os.Args)-1] {
				break
			}
			if v > os.Args[i+1] {
				os.Args[i], os.Args[i+1] = os.Args[i+1], os.Args[i]
			}
		}
	}
	for _, word := range os.Args[1:] {
		for _, letters := range word {
			z01.PrintRune(letters)
		}
		z01.PrintRune('\n')
	}
}
