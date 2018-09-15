package jianzhioffer

import (
	"testing"
)

func mergeList(one, two *NodeList) *NodeList {
	var root, cur, other, tmp *NodeList
	if one.Val <= two.Val {
		root = one
		cur = one
		other = two
	} else {
		root = two
		cur = two
		other = one
	}
	for cur.Next != nil {
		if cur.Next.Val <= other.Val {
			cur = cur.Next
			continue
		}
		tmp = cur.Next
		cur.Next = other
		other = tmp
	}
	cur.Next = other
	return root
}

func TestMergeList(t *testing.T) {
	l1 := buildList([]int{1, 3, 5, 7, 9})
	l2 := buildList([]int{0, 2, 4, 6, 8})
	res := mergeList(l1, l2)
	printList(res)
}

/*
题意：
给两个链表，都是排好序的，现在要将两个链表合并，且合并完之后的链表也是排好序的
思路:
这其实是归并排序的第二步，把两个有序的数组合并成一个有序的数组
*/
