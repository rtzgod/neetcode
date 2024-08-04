package bst

import "fmt"

func DFSInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	DFSInOrder(root.Left)
	fmt.Print(root.Val, "->")
	DFSInOrder(root.Right)
}
func DFSReverseInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	DFSReverseInOrder(root.Right)
	fmt.Print(root.Val, "->")
	DFSReverseInOrder(root.Left)
}
func DFSPreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, "->")
	DFSPreOrder(root.Left)
	DFSPreOrder(root.Right)
}
func DFSPostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	DFSPostOrder(root.Left)
	DFSPostOrder(root.Right)
	fmt.Print(root.Val, "->")
}
