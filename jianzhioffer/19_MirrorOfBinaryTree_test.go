package jianzhioffer

import (
	"fmt"
	"testing"
)

func mirror(root *TreeNode) {
	if root == nil {
		return
	}
	mirror(root.Left)
	mirror(root.Right)
	root.Left, root.Right = root.Right, root.Left
}

func TestMirror(t *testing.T) {
	root := buildTestTree()
	fmt.Println(inOrder(root))
	fmt.Println(preOrder(root))
	mirror(root)
	fmt.Println(inOrder(root))
	fmt.Println(preOrder(root))
}

/*
题意：
二叉树的镜像，左右子树交换且左右子树也镜像
思路：
无脑递归
*/
