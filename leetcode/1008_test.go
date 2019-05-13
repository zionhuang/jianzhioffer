package leetcode

import (
	"fmt"
	"testing"
)

func bstFromPreOrder(preOrder []int) *TreeNode {
	l := len(preOrder)
	if l == 0 {
		return nil
	}
	root := &TreeNode{Val: preOrder[0]}
	m := l
	for i:=1; i<l; i++ {
		if preOrder[i] > preOrder[0] {
			m = i
			break
		}
	}
	root.Left = bstFromPreOrder(preOrder[1:m])
	root.Right = bstFromPreOrder(preOrder[m:])
	return root
}

func Test1007(t *testing.T) {
	root := bstFromPreOrder([]int{8,5,1,7,10,12})
	fmt.Println(root.Left.Val)
}