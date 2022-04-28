package piscine

func ForEach(f func(int), a []int) {
	for _, letter := range a {
		f(letter)
	}
}
