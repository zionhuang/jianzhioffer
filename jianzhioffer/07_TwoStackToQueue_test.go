package jianzhioffer

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"testing"
)

type MyQueue struct {
	InStack  stack.Stack
	OutStack stack.Stack
}

func (q *MyQueue) Push(val interface{}) {
	q.InStack.Push(val)
}

func (q *MyQueue) Pop() interface{} {
	if q.OutStack.Len() > 0 {
		return q.OutStack.Pop()
	}
	for q.InStack.Len() > 0 {
		q.OutStack.Push(q.InStack.Pop())
	}
	return q.OutStack.Pop()
}

func (q *MyQueue) Len() int {
	return q.InStack.Len() + q.OutStack.Len()
}

func TestConvert(t *testing.T) {
	q := MyQueue{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	fmt.Println(q.Pop())
	fmt.Println(q)
}
