package jianzhioffer

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var left, right []int
	if root.Left != nil {
		left = inOrder(root.Left)
	}
	if root.Right != nil {
		right = inOrder(root.Right)
	}
	left = append(left, root.Val)
	left = append(left, right...)
	return left
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res, left, right []int
	res = append(res, root.Val)
	if root.Left != nil {
		left = preOrder(root.Left)
	}
	if root.Right != nil {
		right = preOrder(root.Right)
	}
	res = append(res, left...)
	res = append(res, right...)
	return res
}

func buildTestTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Right.Left = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{6, nil, nil}
	return root
}
