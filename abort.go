package piscine

func Abort(a, b, c, d, e int) int {
	var array []int
	array = append(array, a, b, c, d, e)
	for c := 0; c < 6; c++ {
		for i := 1; i < len(array); i++ {
			if array[i-1] > array[i] {
				array[i-1], array[i] = array[i], array[i-1]
			}
		}
	}
	return array[2]
}
