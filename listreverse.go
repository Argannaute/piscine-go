package piscine

func ListReverse(l *List) {
	ite := l.Head
	counter := -1
	if l.Head == nil {
		return
	}
	for ite != nil {
		ListPushFront(l, ite.Data)
		ite = ite.Next
		counter++
	}
	iterate := l.Head
	for counter != 0 {
		iterate = iterate.Next
		counter--
	}
	iterate.Next = nil
	l.Tail = iterate
}
