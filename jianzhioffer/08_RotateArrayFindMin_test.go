package jianzhioffer

import (
	"fmt"
	"testing"
)

func findMin(arr []int) int {
	left := 0
	right := len(arr) - 1
	if right == 0 {
		return arr[0]
	}
	mid := left
	for arr[left] >= arr[right] {
		mid = (left + right) / 2
		if arr[mid] < arr[left] {
			right = mid
			continue
		}
		if arr[mid] > arr[right] {
			left = mid + 1
			continue
		}
		for i := 0; i < right+1; i++ {
			if arr[i] < arr[mid] {
				mid = i
			}
		}
		break

	}
	return arr[mid]
}

func TestFindMin(t *testing.T) {
	fmt.Println(findMin([]int{8, 1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(findMin([]int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(findMin([]int{1, 2, 2, 2, 2, 0, 1, 1}))
	fmt.Println(findMin([]int{3, 4, 5, 6, 7, 1, 1, 1}))
	fmt.Println(findMin([]int{2, 2, 2, 1, 1, 1, 1, 1}))
	fmt.Println(findMin([]int{1, 0, 1, 1, 1, 1, 1, 1}))

}

/*
题意：
给定一个一维数组，是一个递增数组的旋转，所谓旋转就是：1，2，3，4，5 变成 3，4，5，1，2
设计算法，高效的找到旋转数组的最小值
思路：
1. 一次遍历，O(n)
2. 二分法，尽可能在一次判断后可以丢掉一半
*/
