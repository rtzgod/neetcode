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

func Remove(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val > root.Val {
		root.Right = Remove(root.Right, val)
	} else if val < root.Val {
		root.Left = Remove(root.Left, val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			minNode := minValueNode(root.Right)
			root.Val = minNode.Val
			root.Right = Remove(root.Right, minNode.Val)
		}
	}
	return root
}

func minValueNode(root *TreeNode) *TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}
