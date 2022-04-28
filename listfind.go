package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l.Head == nil {
		return nil
	}
	ite := l.Head
	for ite.Next != nil {
		if CompStr(ref, ite.Data) {
			return &ite.Data
		} else {
			ite = ite.Next
		}
	}
	return nil
}
