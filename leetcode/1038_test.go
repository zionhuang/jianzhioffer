package leetcode

import "testing"

func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Right != nil {
		right := dfs(root.Right, 0)
		root.Val += right
	}
	if root.Left == nil {
		return root
	}
	dfs(root.Left, root.Val)
	return root
}

func dfs(root *TreeNode, input int) int {
	if root.Right != nil {
		right := dfs(root.Right, input)
		root.Val += right
	}
	if root.Left == nil {
		return root.Val
	}
	return dfs(root.Left, root.Val)
}

func TestBstToGst(t *testing.T) {

}