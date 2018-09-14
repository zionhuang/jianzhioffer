package jianzhioffer

import "fmt"

// 链表节点
type NodeList struct {
	Val  int
	Next *NodeList
}

// 根据数组构造链表
func buildList(arr []int) *NodeList {
	l := len(arr)
	if l == 0 {
		return nil
	}
	head := &NodeList{Val: arr[0], Next: nil}
	head.Next = buildList(arr[1:])
	return head
}

// 打印链表
func printList(head *NodeList) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}
