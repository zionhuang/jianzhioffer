package jianzhioffer

import (
	"fmt"
	"testing"
)

func power(base float64, exponent int) float64 {
	// 如果base就是1，不管exponent是多少，结果都是1
	if base == 1.0 {
		return 1.0
	}
	// exponent正负的问题
	flag := 1
	if exponent < 0 {
		// 如果exponent是负数，先将exponent转成正数来处理
		flag = -1
		exponent = -exponent
	}
	var res float64
	switch exponent {
	case 0:
		res = 1
	case 1:
		res = base
	default:
		// 如果exponent大于1，用二分法来处理
		half := power(base, exponent/2)
		mod := exponent % 2
		if mod == 0 {
			res = half * half
		} else {
			// 如果exponent是奇数，二分之后要多乘一个base
			res = half * half * base
		}
	}
	if flag == -1 {
		// exponent是负数的情况，res要倒过来
		res = 1 / res
	}
	return res
}

func TestPower(t *testing.T) {
	fmt.Println(power(0.5, -3))
}

/*
题意：
在不是用库函数的情况下，自己实现求一个数base的exponent次方
思路：
这题很简单，重要是的是把exponent的正数负数零值得情况考虑完整就可以了
*/
