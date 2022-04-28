package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	ans := l
	for pos != 0 {
		if ans.Next == nil {
			return nil
		}
		ans = ans.Next
		pos--
	}
	return ans
}
