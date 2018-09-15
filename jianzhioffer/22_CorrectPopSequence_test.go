package jianzhioffer

import (
	"fmt"
	"github.com/badgerodon/collections/stack"
	"testing"
)

func correctSequence(push, pop []int) bool {
	l1 := len(push)
	l2 := len(pop)
	if l1 != l2 {
		// 如果进出栈的数量都不同，返回false
		return false
	}
	if l1 == 0 {
		// 都是空，返回true
		return true
	}
	in := stack.New()
	i, j := 0, 0 // i是入栈顺序指针 j是出栈顺序指针
	for i < l1 {
		in.Push(push[i])
		if in.Peek() == pop[j] {
			// 如果栈顶元素和当前的出栈元素相等就出栈
			in.Pop()
			// 出栈指针后移
			j++
			for in.Len() > 0 && in.Peek() == pop[j] {
				// 如果入栈还有元素，继续比，有机会就出栈
				in.Pop()
				j++
			}
		}
		// 如果当前入栈元素和出栈元素不等，就继续压栈
		i++
	}
	if j < l2 {
		// 如果是正确的出栈顺序，j的长度为l2，如果j<l2，一定是出栈顺序不对，导致出栈不干净
		return false
	}
	return true
}

func TestCorrectSequence(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{5, 4, 3, 2, 1}
	fmt.Println(correctSequence(a, b))
}

/*
题意：
给定一个栈的入栈顺序，再给一个出栈顺序，判断这个出栈顺序是否是正确的
思路：
用一个栈来进行辅助入栈
*/
