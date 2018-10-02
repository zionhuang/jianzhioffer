package leetcode

import (
	"testing"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var slow, fast, pre *ListNode
	slow = head
	fast = head
	for i:=1; i<n; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	if pre == nil {
		return head.Next
	}
	pre.Next = slow.Next
	return head
}

func TestRemoveNthFromEnd(t *testing.T)  {
	values := []int{1, 2, 3, 4, 5}
	head := buildList(values)
	printList(head)
	newHead := removeNthFromEnd(head, 2)
	printList(newHead)
}

/*
题意: 删除一个链表中的倒数第k个元素
思路: 这题和剑指offer上那题找倒数第k个数的题相似，多了删除的操作，记录倒数第k个数的前驱，用来做删除操作即可
 */