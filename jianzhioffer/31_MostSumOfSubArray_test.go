package jianzhioffer

import (
	"testing"
	"fmt"
)

func mostSubSum(arr []int) (res int) {
	dp := make([]int, len(arr)+1)
	for i:=0; i<len(arr); i++ {
		tmp := dp[i] + arr[i]
		if tmp > 0 {
			dp[i+1] = tmp
		} else {
			dp[i+1] = 0
		}
	}
	for i:=1; i<=len(arr); i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return
}

func TestMostSubSum(t *testing.T) {
	arr := []int{1, -2, 3, 10, -4, 7, 2, -5}
	res := mostSubSum(arr)
	fmt.Println(res)
}

/*
题意:
找到和最大的子序列
思路:
 */