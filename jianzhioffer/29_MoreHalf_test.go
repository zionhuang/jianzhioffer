package jianzhioffer

import (
	"fmt"
	"math"
	"testing"
)

// 字典
func moreHalfByDic(arr []int) (res int) {
	if arr == nil {
		return
	}
	if len(arr) <= 2 {
		res = arr[0]
		return
	}
	var dic = make(map[int]int)
	var most = math.MinInt64
	for _, k := range arr {
		dic[k]++
		if dic[k] > most {
			res = k
			most = dic[k]
		}
	}
	return
}

func TestMoreHalfByDic(t *testing.T) {
	values := []int{1, 2, 3, 2, 2}
	res := moreHalfByDic(values)
	fmt.Println(res)
}

// 通过抵消
func moreHalfByAddSub(arr []int) (res int) {
	if arr == nil {
		return
	}
	if len(arr) <= 2 {
		res = arr[0]
		return
	}
	var cur int
	var cnt int
	cur = arr[0]
	cnt = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == cur {
			cnt++
			continue
		}
		cnt--
		if cnt > 0 {
			continue
		}
		cur = arr[i]
		cnt = 1
	}
	return cur
}

func TestMoreHalfByAddSub(t *testing.T) {
	values := []int{3, 3, 3, 2, 2}
	res := moreHalfByAddSub(values)
	fmt.Println(res)
}

/*
题意:
给定一个数组，里面有一个数超过总数的一半，找出这个数
思路:
思路1. 用字典计数，并记录当前最多的那个数，最后输出 O(n) 最多有 n/2 的空间复杂度
思路2. 通过抵消的方式
*/
