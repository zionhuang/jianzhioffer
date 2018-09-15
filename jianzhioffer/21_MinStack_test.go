package jianzhioffer

import (
	"fmt"
	"github.com/badgerodon/collections/stack"
	"testing"
)

type MinStack struct {
	Val *stack.Stack
	Min *stack.Stack
}

func (m *MinStack) Init() {
	m.Val = stack.New()
	m.Min = stack.New()
}

func (m *MinStack) push(val interface{}) {
	m.Val.Push(val)
	if m.len() == 0 || val.(int) < m.min() {
		m.Min.Push(val)
	} else {
		m.Min.Push(m.min())
	}
}

func (m *MinStack) min() int {
	return m.Min.Peek().(int)
}

func (m *MinStack) pop() int {
	m.Min.Pop()
	return m.Val.Pop().(int)
}

func (m *MinStack) len() int {
	return m.Min.Len()
}

func TestMinStack(t *testing.T) {
	m := &MinStack{nil, nil}
	m.Init()
	a := []int{1, 2, 0, 4, 5, 6, -1}
	for _, v := range a {
		m.push(v)
	}
	for m.len() > 0 {
		fmt.Println(m.min())
		fmt.Println(m.pop())
	}

}

/*
题意：
实现一个可以快速获取最小值得栈
思路：
用两个栈，一个栈作为主栈，功能就是普通的栈，先进后出，另一个栈用来存储当前的最小值，
*/
