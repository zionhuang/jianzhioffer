package leetcode

import (
	"testing"
)

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var cur *ListNode
	cur = lists[0]
	for i:=1; i<len(lists); i++ {
		cur = mergeList(cur, lists[i])
	}
	return cur
}

func mergeList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head, this, other, tmp *ListNode
	if list1.Val > list2.Val {
		head = list2
		this = list2
		other = list1
	} else {
		head = list1
		this = list1
		other = list2
	}
	for this.Next != nil {
		if this.Next.Val > other.Val {
			tmp = this.Next
			this.Next = other
			other = tmp
			this = this.Next
		} else {
			this = this.Next
		}
	}
	this.Next = other
	return head
}

func TestMergeKLists(t *testing.T)  {
	var values [][]int
	values = [][]int{
		{1,3,5,7,9},
		{1,2,4,6,7},
		{10, 12, 30},
		{},
		{3,3,3,3,3,3},
	}
	var lists[]*ListNode
	for i := range values {
		tmp := buildList(values[i])
		lists = append(lists, tmp)
	}
	res := mergeKLists(lists)
	printList(res)
}

/*
题意:
把k个有序的链表合并，且合并之后的链表也是有序的
思路:
这题是合并两个有序链表的延伸，
思路1，一个一个的去合并就好了
思路2，利用优先队列
 */