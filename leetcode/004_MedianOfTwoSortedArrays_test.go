package leetcode

import (
	"fmt"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var flag bool
	m := len(nums1)
	n := len(nums2)
	total := m + n
	if total&1 == 1 {
		flag = true
	}
	mid := total / 2
	i := 0
	j := 0
	var pre int
	for i < m && j < n && mid > 0 {
		if nums1[i] < nums2[j] {
			pre = nums1[i]
			i++
			mid--
		} else {
			pre = nums2[j]
			j++
			mid--
		}
		//fmt.Printf("i:%v, j:%v, mid:%v\n", i, j, mid)
	}
	//fmt.Printf("i:%v, j:%v, mid:%v\n", i, j, mid)
	// nums1小且少 先越界的情况
	if i == m {
		for mid > 0 {
			pre = nums2[j]
			j++
			mid--
		}
		if flag {
			//fmt.Printf("i:%v, j:%v, mid:%v\n", i, j, mid)
			return float64(nums2[j])
		} else {
			return float64(pre+nums2[j]) / 2
		}
	}

	// nums2小且少 先越界的情况
	if j == n {
		for mid > 0 {
			pre = nums1[i]
			i++
			mid--
		}
		if flag {
			return float64(nums1[i])
		} else {
			return float64(pre+nums1[i]) / 2
		}
	}

	//
	var cur int
	if nums1[i] < nums2[j] {
		cur = nums1[i]
	} else {
		cur = nums2[j]
	}
	if flag {
		return float64(cur)
	}
	return float64(pre+cur) / 2
}

func TestFindMedianSortedArrays(t *testing.T) {
	// 已经约定了 两个数组不会都是空
	nums11 := []int{}
	nums12 := []int{1}
	res1 := findMedianSortedArrays(nums11, nums12)
	fmt.Println(res1) // 1

	nums11 = []int{2, 3}
	nums12 = []int{}
	res1 = findMedianSortedArrays(nums11, nums12)
	fmt.Println(res1) // 2.5

	nums21 := []int{1, 3}
	nums22 := []int{2, 4, 5, 6, 7, 8}
	res2 := findMedianSortedArrays(nums21, nums22)
	fmt.Println(res2) // 4.5

	nums31 := []int{1, 3}
	nums32 := []int{2, 4, 5, 6, 7, 8, 9}
	res3 := findMedianSortedArrays(nums31, nums32)
	fmt.Println(res3) // 5

	nums41 := []int{1, 2, 3, 4}
	nums42 := []int{5, 6, 7, 8}
	res4 := findMedianSortedArrays(nums41, nums42)
	fmt.Println(res4) // 5
}

/*
题意：
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
You may assume nums1 and nums2 cannot be both empty.
两个已经排好序的数组求中位数
思路：
我之前用Python写过这题，当时就是把前mid+1个数装进了一个list，要么输出最后一个数，要么是最后两个数的平均数
这次看到题首先想到的就是利用归并排序的第二步归并的方式找到中间的那个或那两个数。
但是这次写得很慢，各种条件考虑的不清楚
*/
