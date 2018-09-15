package jianzhioffer

import (
	"fmt"
	"testing"
)

func count(n int) int {
	count := 0
	for n != 0 {
		count++
		n = n & (n - 1)
	}
	return count
}

func TestCount(t *testing.T) {
	fmt.Println(count(7))
}

/*
题意：
计算一个十进制数的二进制数中的1的个数
思路：
这题是考察位运算
1. 把十进制数转化成二进制数，然后遍历二进制中1的个数。效率最低
2. n和1相与，如果结果是1，count加1，然后n右移一位。这个方法从n的二进制的最低位去判断是否是1，效率提升，判断的次数与二进制的位数相同
3. n和n-1相与，会把n的二进制数的最右边的1消掉，一直消掉n等于。效率最高，n的二进制中有几个1就执行几次。
*/
