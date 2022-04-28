package piscine

func ConcatParams(args []string) string {
	answer := ""
	first := false
	for i := 0; i < len(args); i++ {
		if first {
			answer += "\n"
		}
		answer += args[i]
		first = true
	}
	return answer
}
