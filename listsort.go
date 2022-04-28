package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return l
	}
	ite := l
	for i := 0; i < 200; i++ {
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
