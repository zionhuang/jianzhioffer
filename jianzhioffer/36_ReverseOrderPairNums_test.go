package jianzhioffer

import (
	"fmt"
	"testing"
)

var cnt int

func TestReverseOrderPairNum(t *testing.T) {
	arr := []int{7, 5, 6, 3}
	res := reverseOrderPairNum(arr)
	fmt.Println(res)
	fmt.Println(cnt)
}

func reverseOrderPairNum(arr []int) []int {
	l := len(arr)
	mid := l / 2
	if mid == 0 {
		return arr
	}
	var res []int
	left := reverseOrderPairNum(arr[:mid])
	right := reverseOrderPairNum(arr[mid:])
	res = merge(left, right)
	return res
}

func merge(a, b []int) (res []int) {
	l1 := len(a)
	l2 := len(b)
	i := 0
	j := 0
	for i <= l1-1 && j <= l2-1 {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
			continue
		}
		cnt = cnt + l1 - i
		res = append(res, b[j])
		j++
	}
	if i == l1 {
		res = append(res, b[j:]...)
		return
	}
	res = append(res, a[i:]...)
	return
}

func TestMerge(t *testing.T) {
	a := []int{2, 3}
	b := []int{0, 1}
	res := merge(b, a)
	fmt.Println(res)
	fmt.Println(cnt)
}

/*
题意:
给定一个数组，找出有多少对逆序对
思路:
[7 5 6 3] (7 5) (7 6) (7 3) (5 3) (6 3) 一共5对
用冒泡的思路可以，但是时间复杂度是 O(n^2)
用归并可以，时间复杂度是 O(nlogn)
为什么冒泡和归并可以？而快排不可以？我想应给是因为稳定性，快排在一次交换后可能会改变原本的逆序对
*/
