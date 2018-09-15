package jianzhioffer

import (
	"fmt"
	"testing"
)

func changeOddEven(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]%2 == 1 {
			i++
			for i < j && arr[i]%2 == 1 {
				i++
			}
		}
		if arr[j]%2 == 0 {
			j--
			for i < j && arr[j]%2 == 0 {
				j--
			}
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
}

func TestChangeOddEven(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	changeOddEven(a)
	fmt.Println(a)
}

/*
题意：
给定一个数组arr， 把所有奇数移到偶数前面
思路：
如果这题还要保证奇数们、偶数们的相对位置不变会复杂一点，这题用首尾双指针就可以行了
*/
