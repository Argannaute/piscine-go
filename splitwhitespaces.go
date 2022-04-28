package piscine

func SplitWhiteSpaces(s string) []string {
	answer := []string{}
	var word string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' || s[i] == '\t' {
			if word != "" && word != " " {
				answer = append(answer, word)
				word = ""
			}
		} else {
			word += string(s[i])
		}
		if i == len(s)-1 {
			if word != "" {
				answer = append(answer, word)
			}
		}
	}
	return answer
}
