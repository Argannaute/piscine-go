package main

import (
	"os"
)

func main() {
	if os.Args[2] != "+" && os.Args[2] != "-" && os.Args[2] != "/" && os.Args[2] != "*" && os.Args[2] != "%" {
		os.Exit(0)
	}
	nbrArr := []rune{}
	res := 0
	nbr1 := Atoi(os.Args[1])
	nbr2 := Atoi(os.Args[3])
	divi := "No divsion by 0"
	mod := "No modulo by 0"
	for i := 0; i < len(os.Args[1]); i++ {
		args := os.Args[1]
		nbrArr = append(nbrArr, rune(args[i]))
	}
	for i := 0; i < len(os.Args[3]); i++ {
		args := os.Args[3]
		nbrArr = append(nbrArr, rune(args[i]))
	}
	for _, val := range nbrArr {
		if val < 48 || val > 57 {
			os.Exit(0)
		}
	}
	if nbr1 == 0 && os.Args[2] == "%" || nbr2 == 0 && os.Args[2] == "%" {
		os.Stdout.WriteString(mod)
		os.Exit(0)
	}
	if nbr1 == 0 && os.Args[2] == "/" || nbr2 == 0 && os.Args[2] == "/" {
		os.Stdout.WriteString(divi)
		os.Exit(0)
	}
	if nbr1 == 9223372036854775807 || nbr1 == -9223372036854775808 {
		os.Exit(0)
	}
	if nbr2 == 9223372036854775807 || nbr2 == -9223372036854775808 {
		os.Exit(0)
	}
	if os.Args[2] == "+" {
		res = nbr1 + nbr2
	}
	if os.Args[2] == "-" {
		res = nbr1 - nbr2
	}
	if os.Args[2] == "/" {
		res = nbr1 / nbr2
	}
	if os.Args[2] == "%" {
		res = nbr1 % nbr2
	}
	if os.Args[2] == "*" {
		res = nbr1 * nbr2
	}
	for _, val := range string(res) {
		os.Stdout.WriteString(stringVarInside)
	}
}

func Atoi(s string) int {
	result := 0
	sign := 1
	if s == "" || s == "0" {
		return 0
	}
	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + (int(ch) - '0')
	}
	return (result * sign)
}
