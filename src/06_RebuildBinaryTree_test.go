package src

import (
	"fmt"
	"testing"
)

func rebuild(preOrder, inOrder []int) *TreeNode {
	lPre := len(preOrder)
	lIn := len(inOrder)
	if (lPre == 0) || (lIn == 0) || (lPre != lIn) {
		return nil
	}
	a := preOrder[0]
	head := &TreeNode{a, nil, nil}
	point := 0
	for i, v := range inOrder {
		if v == a {
			point = i
		}
	}
	head.Left = rebuild(preOrder[1:point+1], inOrder[0:point])
	if point+1 >= lIn {
		return head
	}
	head.Right = rebuild(preOrder[point+1:], inOrder[point+1:])
	return head
}

func TestBuildBinaryTree(t *testing.T) {
	pre := []int{4, 1, 2, 3, 5, 7}
	in := []int{2, 1, 4, 5, 3, 7}
	binaryTree := rebuild(pre, in)
	fmt.Println(binaryTree)
	fmt.Println(binaryTree.Left, binaryTree.Right)
	fmt.Println(binaryTree.Left.Left, binaryTree.Left.Right, binaryTree.Right.Left, binaryTree.Right.Right)
}
