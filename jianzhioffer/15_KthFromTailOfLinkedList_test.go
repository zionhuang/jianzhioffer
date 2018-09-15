package jianzhioffer

import (
	"fmt"
	"testing"
)

func findKthTail(head *NodeList, k int) int {
	slow := head
	quick := head
	for i := 1; i < k; i++ {
		if quick.Next == nil {
			// k比链表长度长
			return 404
		}
		quick = quick.Next
	}
	for quick.Next != nil {
		quick = quick.Next
		slow = slow.Next
	}
	return slow.Val
}

func TestFindKthTail(t *testing.T) {
	ll := buildList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(findKthTail(ll, 9))
	fmt.Println(findKthTail(ll, 10))
	fmt.Println(findKthTail(ll, 4))
}

/*
题意：
给定一个链表，找到倒数第k个节点的数。
思路：
1. 两次遍历，第一次记录长度
2. 遍历一次，用两个指针，让一个指针先走k-1步，那么先走的指针走到末尾的时候，后走的指针走了(length-k+1步)，那么和最后一个相差(k-1)步，即倒数第k个节点
*/
