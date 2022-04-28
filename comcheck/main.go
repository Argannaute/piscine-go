package main

import (
	"fmt"
	"os"
)

func main() {
	str := "Alert!!!"
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i] == "01" || os.Args[i] == "galaxy" || os.Args[i] == "galaxy 01" {
			fmt.Println(str)
			os.Exit(0)
		}
	}
}
