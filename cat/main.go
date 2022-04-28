package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		b, _ := ioutil.ReadAll(os.Stdin)
		printStr(string(b))
		os.Exit(0)
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		error := "ERROR: "
		error += err.Error()
		printStr(error)
		z01.PrintRune('\n')
		os.Exit(1)
	} else {
		printStr(string(content))
	}
	if len(os.Args) == 3 {
		content, err := os.ReadFile(os.Args[2])
		if err != nil {
			error := "ERROR: "
			error += err.Error()
			printStr(error)
			z01.PrintRune('\n')
			os.Exit(1)
		} else {
			printStr(string(content))
		}
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
