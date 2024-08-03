package bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Search(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > val {
		return Search(root.Left, val)
	} else if root.Val < val {
		return Search(root.Right, val)
	} else {
		return root
	}
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if val < root.Val {
		root.Left = Insert(root.Left, val)
	} else if val > root.Val {
		root.Right = Insert(root.Right, val)
	}

	return root
}
