package jianzhioffer

import (
	"fmt"
	"testing"
)

var times int

func find(arr [][]int, value int) bool {
	r := len(arr)
	if r == 0 {
		return false
	}
	c := len(arr[0])
	if c == 0 {
		return false
	}
	i := r - 1
	j := 0
	for i >= 0 && j < c {
		times++
		if arr[i][j] == value {
			return true
		}
		if arr[i][j] > value {
			i--
		} else {
			j++
		}
	}
	return false
}

func TestFind(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
		{5, 6, 7, 8, 9},
	}
	fmt.Println(find(arr, 3))
	fmt.Println(times)
}
