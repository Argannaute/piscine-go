package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	ite := l.Head
	if l.Head.Data == data_ref {
		l.Head = ite.Next
	}
	for ite.Next != nil {
		if ite.Next.Data == data_ref {
			ite.Next = ite.Next.Next
		} else {
			ite = ite.Next
		}
	}
}
