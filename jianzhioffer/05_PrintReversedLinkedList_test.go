package jianzhioffer

import (
	"testing"
)

// 单链表翻转
func reverse(head *NodeList) *NodeList {
	if head == nil {
		return nil
	}
	var cur, pre *NodeList
	for head.Next != nil {
		cur = head.Next
		head.Next = pre
		pre = head
		head = cur
	}
	head.Next = pre
	return head
}

// 逆序打印链表
func reversePrint(head *NodeList) {
	printList(reverse(head))
}

func TestBuild(t *testing.T) {
	arr := []int{1, 2, 3, 8, 9, 3, 5, 7, 4, 5}
	list := buildList(arr)
	reversePrint(list)
}
