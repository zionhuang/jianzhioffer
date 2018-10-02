package leetcode

import "testing"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k < 2 {
		return head
	}
	var cur *ListNode
	cur = head
	i := k
	for i > 0 && cur != nil {
		cur = cur.Next
		i--
	}
	if i > 0 {
		return head
	}
	cur = reverseKGroup(cur, k)
	head = reverse(head, k, cur)
	return head
}

func reverse(head *ListNode, k int, tail *ListNode) *ListNode {
	var pre, cur, tmp *ListNode
	pre = tail
	cur = head
	for i:=0; i<k; i++ {
		tmp = cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func TestReverseKGroup(t *testing.T)  {
	values := []int{1,2,3,4,5,6,7,8,9,10,11}
	ln := buildList(values)
	printList(ln)
	res := reverseKGroup(ln, 4)
	printList(res)
}

/*
题意:
给定一个链表，以K为窗口进行翻转，比如链表 1->2->3->4->5->6->7->8->9->10->11 以k=4为窗口进行翻转 4->3->2->1->8->7->6->5->9->10->11
思路:
链表翻转的变体，可以用递归的思路，整个翻转过程可以分成：1.翻转前k个 2.翻转后面的所有
 */