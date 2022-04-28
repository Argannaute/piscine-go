package piscine

func ListSize(l *List) int {
	count := 1
	n := &NodeL{}

	if l.Head == nil {
		l.Head = n
	} else {
		bana := l.Head
		for bana.Next != nil {
			bana = bana.Next
			count++
		}
		bana.Next = n
	}
	if count == 1 {
		count = 0
	}
	return count
}
