package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil && l2.Head == nil {
		return
	}
	ite1 := l1.Head
	ite2 := l2.Head
	if ite1 == nil {
		l1.Head = l2.Head
		return
	}
	ite1 = l1.Head
	for ite2 != nil {
		if ite1.Next == nil {
			ListPushBack(l1, ite2.Data)
			ite2 = ite2.Next
			ite1 = ite1.Next
		} else {
			ite1 = ite1.Next
		}
	}
}
