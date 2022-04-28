package piscine

func ListLast(l *List) interface{} {
	iterator := l.Head
	if l.Head == nil {
		return nil
	} else {
		for iterator.Next != nil {
			iterator = iterator.Next
		}
	}
	return iterator.Data
}
