package piscine

import "fmt"

func Split(s, sep string) []string {
	answer := []string{}
	var word string
	length := len(sep)
	fmt.Println(sep[0:length])

	for i := 0; i < len(s); i++ {
		if s[i:i+length] == sep[0:] {
			fmt.Println("here")
			if word != "" && word != " " {
				answer = append(answer, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
	}
	return answer
}
