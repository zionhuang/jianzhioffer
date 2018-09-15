package jianzhioffer

import (
	"fmt"
	"testing"
)

func fabonacci(n int) int {
	if n <= 1 {
		return n
	}
	pre := 0
	cur := 1
	for n > 1 {
		tmp := pre + cur
		pre = cur
		cur = tmp
		n--
	}
	return cur
}

func TestFabonacci(t *testing.T) {
	fmt.Println(fabonacci(1))
	fmt.Println(fabonacci(2))
	fmt.Println(fabonacci(3))
	fmt.Println(fabonacci(7))
	fmt.Println(fabonacci(20))
	fmt.Println(fabonacci(30))
	fmt.Println(fabonacci(100))
}

/*
题意：fabonacci 数， 不用解释。。。
*/
