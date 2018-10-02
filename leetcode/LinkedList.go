package leetcode

import "fmt"

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 根据数组构造链表
func buildList(arr []int) *ListNode {
	l := len(arr)
	if l == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0], Next: nil}
	head.Next = buildList(arr[1:])
	return head
}

// 打印链表
func printList(head *ListNode) {
	cur := head
	res := ""
	flag := 1
	for cur != nil {
		if flag == 1 {
			res += fmt.Sprintf("%v", cur.Val)
		} else {
			res += fmt.Sprintf("->%v", cur.Val)
		}
		flag++
		cur = cur.Next
	}
	fmt.Println(res)
}
