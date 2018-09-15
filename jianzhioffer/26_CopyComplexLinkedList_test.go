package jianzhioffer

import (
	"fmt"
	"testing"
)

type ComplexListNode struct {
	Val     int
	Next    *ComplexListNode
	Sibling *ComplexListNode
}

func copyComplex(c *ComplexListNode) *ComplexListNode {
	if c == nil {
		return nil
	}
	// 第一遍，向后复制 1->2->3->4 ===> 1->1->2->2->3->3->4->4
	// 复制的 sibling 暂时为 nil
	var cur, tmp *ComplexListNode
	cur = c
	for cur != nil {
		tmp = cur.Next
		cur.Next = &ComplexListNode{cur.Val, tmp, nil}
		cur = tmp
	}
	// 第二遍，把复制节点的 sibling 赋上值
	// 这里很舒服，新节点的sibling的地址要么和pre的sibling一起是nil，要么是pre的sibling的next
	cur = c
	for cur != nil {
		if cur.Sibling != nil {
			cur.Next.Sibling = cur.Sibling.Next
		}
		cur = cur.Next.Next
	}
	// 把链表拆成两个,返回新生成的链表
	cur = c
	var newHead, even *ComplexListNode
	newHead = cur.Next
	even = newHead
	for cur != nil {
		cur.Next = cur.Next.Next
		if cur.Next != nil {
			even.Next = cur.Next.Next
		} else {
			even.Next = nil
		}
		cur = cur.Next
		even = even.Next
	}
	return newHead
}

func buildComplex(a []int) *ComplexListNode {
	l := len(a)
	if l == 0 {
		return nil
	}
	var head *ComplexListNode
	head = &ComplexListNode{a[0], nil, nil}
	if l == 1 {
		return head
	}
	head.Next = buildComplex(a[1:])
	return head
}

func printComplex(a *ComplexListNode) {
	cur := a
	for cur != nil {
		fmt.Printf("%p  %+v\n", cur, cur)
		cur = cur.Next
	}
}

func TestCopyComplex(t *testing.T) {
	a := []int{1, 2, 3, 4}
	l := buildComplex(a)
	l.Sibling = l.Next.Next
	l.Next.Sibling = l
	l.Next.Next.Next.Sibling = l.Next
	printComplex(l)
	c := copyComplex(l)
	printComplex(c)
}
