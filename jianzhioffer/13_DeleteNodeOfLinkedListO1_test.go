package jianzhioffer

import (
	"fmt"
	"testing"
)

func deleteNode(root, node *NodeList) {
	// node 一定要在root所在的链表中
	if root == nil || node == nil {
		return
	}
	if node.Next == nil {
		node = nil
		return
	}
	node.Val = node.Next.Val
	//tmp := node.Next
	node.Next = node.Next.Next
	// 这里直接把 node.next 跳过了，在c/c++里面肯定要手动释放空间，go这里好像没法释放，估计只能等gc自己回收了
}

func TestDeleteNode(t *testing.T) {
	ll := buildList([]int{1, 2, 3, 4, 5, 6})
	fmt.Println("删除后")
	printList(ll)
	deleteNode(ll, ll.Next.Next)
	fmt.Println("删除后")
	printList(ll)
}

/*
题意：
以O(1)删除链表中的某个节点
思路：
如果要删除的节点给的是value，或者说要删除第几个节点，以O(1)的时间复杂度删除是不可能的，这里的节点给的是地址，其他的没什么好说的
*/
