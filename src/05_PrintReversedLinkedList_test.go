package src

import (
	"fmt"
	"testing"
)

type NodeList struct {
	Val  int
	Next *NodeList
}

// 单链表逆序
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

// 逆序打印链表
func reversePrint(head *NodeList) {
	printList(reverse(head))
}

func TestBuild(t *testing.T) {
	arr := []int{1, 2, 3, 8, 9, 3, 5, 7, 4, 5}
	list := buildList(arr)
	reversePrint(list)
}
