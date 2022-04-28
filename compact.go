package piscine

func Compact(ptr *[]string) int {
	count := 0
	empty := []string{}
	for i, val := range *ptr {
		if i == 0 {
			*ptr = empty
		}
		if val != "" {
			*ptr = append(*ptr, val)
		}
	}
	count = len(*ptr)
	return count
}
