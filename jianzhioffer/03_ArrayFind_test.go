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

/*
题目：
给定一个二维数组array，二维数组的性质是：从左到右递增，从上到下递增，且没有重复的数字。给一个数n，找一下n是否在array中。
思路：
这题如果直接遍历的话，最差的情况是O(n*m)，关键是找到一个点，如果比它大，就可以丢掉一半，比它小也可以丢掉一半，那么就可以不用比较n*m次了
选择 array[n-1][0] 为起始点，这一点的特性是，上面的数都比它小，右边的数都比它大，如果相等，直接返回，如果大，就往上，如果小，就往右
*/
