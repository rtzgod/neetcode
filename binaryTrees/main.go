package main

import (
	"fmt"

	bst "github.com/rtzgod/neetcode/binaryTrees/BST"
)

func main() {
	root := &bst.TreeNode{Val: 1}
	root.Left = &bst.TreeNode{Val: 0}
	root.Right = &bst.TreeNode{Val: 2}
	bst.Insert(root, 5)
	fmt.Println(bst.Search(root, 5))
}
