package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	ListPushBck(l, data_ref)
	ite := l
	for i := 0; i < 1000; i++ {
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

func ListPushBck(l *NodeI, data int) {
	n := &NodeI{Data: data}
	bana := l
	for bana.Next != nil {
		bana = bana.Next
	}
	bana.Next = n
}
