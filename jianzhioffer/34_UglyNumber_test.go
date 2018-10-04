package jianzhioffer

import (
	"fmt"
	"testing"
)

func TestUglyNumber(t *testing.T) {
	res, un := uglyNumber(1500)
	fmt.Println(res)
	fmt.Println(un)
}

func uglyNumber(size int) (int, []int) {
	// 用来存所有的(size个)丑数
	un := make([]int, size)
	// 四个指针和初始值的确定
	pt2 := 0  // 乘2 当前指针
	pt3 := 0  // 乘2 当前指针
	pt5 := 0  // 乘5 当前指针
	un[0] = 1 // 第一个丑数 1
	pt := 1   // 下一个丑数的指针
	for pt < size {
		min := min(un[pt2]*2, un[pt3]*3, un[pt5]*5)
		un[pt] = min
		for un[pt2]*2 == un[pt] {
			// 如果min是pt2上的数乘2产生的，2的指针要后移
			pt2++
		}
		for un[pt3]*3 == un[pt] {
			// 如果min是pt3上的数乘3产生的，3的指针要后移
			pt3++
		}
		for un[pt5]*5 == un[pt] {
			// 如果min是pt5上的数乘5产生的，5的指针要后移
			pt5++
		}
		// 上面的做法还很好的去了重，pt3指向2的，pt2指向3，pt5指向2的时候，
		// 此时最大丑数是5，下一个出现通过计算可知是 6，pt2 和 pt3 此时都可以得到6
		// 这时 pt2 和 pt3 都会后移
		pt++
	}
	return un[pt-1], un
}

func min(a, b, c int) (min int) {
	if a <= b {
		min = a
	} else {
		min = b
	}
	if c <= min {
		min = c
	}
	return
}

/*
题意:
丑数: 因子只包含 2、、3、5 的数被称为丑数，1 2 3 4 5 6 8 9 10 12 ...
找到从小到大第1500个丑数
思路:
思路一：从1开始找，直到找到第1500个。但是这种方式会导致大量非丑数被计算和判断是否是丑数，效率不高
思路二：因为丑数是任意个 2、3、5 相乘的结果，我们可以用一个数组存之前的丑数，并用之前的丑数乘2、3、5得到新的丑数，并保证从小到大排下去
虽然一开始想到的就是第二种思路，但是代码写不出来，不知道怎么保证有序进行，这里其实用到了4个指针，比常见的双指针问题要难一点，
其中三个指针分别对应 2 3 5 这三个数字的"该乘"位置，所谓"该乘"，就是一轮结束后 pt2*2 pt3*3 pt*5 一定是大于当前最大的丑数，
另一个指针指向新丑数出现的位置，如果已经到了 1500 就结束
*/
