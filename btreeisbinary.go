package piscine

func BTreeIsBinary(root *TreeNode) bool {
	right := true
	left := true
	if root == nil {
		return true
	}
	if root.Right != nil {
		right = BTreeIsBinary(root.Right) && root.Right.Data >= root.Data
	}
	if root.Left != nil {
		left = BTreeIsBinary(root.Left) && root.Left.Data <= root.Data
	}
	return left && right
}
