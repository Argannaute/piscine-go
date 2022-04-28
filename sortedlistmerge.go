package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	l := Listmerge(n1, n2)
	ans := Listsort(l)
	return ans
}

func Listmerge(l1 *NodeI, l2 *NodeI) *NodeI {
	if l1 == nil && l2 == nil {
		return nil
	}
	ite1 := l1
	ite2 := l2
	if ite1 == nil {
		l1 = l2
		return l1
	}
	ite1 = l1
	for ite2 != nil {
		if ite1.Next == nil {
			ListPushbck(l1, ite2.Data)
			ite2 = ite2.Next
			ite1 = ite1.Next
		} else {
			ite1 = ite1.Next
		}
	}
	return l1
}

func Listsort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	ite := l
	for i := 0; i < 500; i++ {
		if ite.Next == nil {
			ite = l
		}
		if ite.Data > ite.Next.Data {
			save := ite.Data
			ite.Data = ite.Next.Data
			ite.Next.Data = save
		}
		ite = ite.Next
	}
	return l
}

func ListPushbck(l *NodeI, data int) {
	n := &NodeI{Data: data}
	bana := l
	for bana.Next != nil {
		bana = bana.Next
	}
	bana.Next = n
}
