package jianzhioffer

import (
	"fmt"
	"strconv"
	"testing"
)

func inc(count []int) {
	i := 0
	tmp := 1
	for tmp == 1 {
		if count[i] < 9 {
			count[i] = count[i] + 1
			break
		}
		count[i] = 0
		tmp = 1
		i++
	}
}

func countToString(count []int) string {
	l := len(count)
	i := l - 2
	for count[i] == 0 {
		i--
	}
	res := ""
	for i >= 0 {
		res += strconv.Itoa(count[i])
		i--
	}
	return res
}

func printMaxNBitNumber(n int) {
	count := make([]int, n+1)
	count[0] = 1
	for count[n] == 0 {
		fmt.Println(countToString(count))
		inc(count)
	}
}

func TestPrintMaxNBitNumber(t *testing.T) {
	printMaxNBitNumber(1)
}

/*
题意：给一个数n，打印从1到最大的n位数的所有数，比如n=5，就要从1打印到99999
思路：
这题咋一看很简单，但是陷阱是大数问题，即：在n很大的情况下，比如n=20，最大的数已经超出了int甚至int64的范围，
这种情况，我们只能用数组去模拟数字
*/
