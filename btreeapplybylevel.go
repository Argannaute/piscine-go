package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	h := BTreeLevelcount(root)
	for i := 0; i < h; i++ {
		applyLevel(root, i, f)
	}
}

func applyLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 0 {
		f(root.Data)
	} else if level > 0 {
		applyLevel(root.Left, level-1, f)
		applyLevel(root.Right, level-1, f)
	}
}

func BTreeLevelcount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := BTreeLevelcount(root.Left)
	right := BTreeLevelcount(root.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func BTreeInsertdata(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		root.Left = BTreeInsertdata(root.Left, data)
		root.Left.Parent = root
	} else if data > root.Data {
		root.Right = BTreeInsertdata(root.Right, data)
		root.Right.Parent = root
	}
	return root
}
