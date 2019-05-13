package leetcode

import (
	"github.com/golang-collections/collections/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: a[0],
	}
	stack.New()
	root.Left = &TreeNode{

	}
	return nil
}