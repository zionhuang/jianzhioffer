package jianzhioffer

import (
	"fmt"
	"sort"
	"strconv"
	"testing"
)

// 字符数组排序的准备工作---实现排序接口
type ArrString []string

func (as ArrString) Len() int {
	return len(as)
}

func (as ArrString) Less(i, j int) bool {
	return less(as[i], as[j])
}

func (as ArrString) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}

// 整数数组转成字符串数组
func intArrayToStringArray(intArr []int) (res ArrString) {
	for _, v := range intArr {
		tmp := strconv.Itoa(v)
		res = append(res, tmp)
	}
	return
}

// less规则
func less(a, b string) bool {
	l1 := len(a)
	l2 := len(b)
	i := 0
	for i < l1 && i < l2 {
		// 1 < 2; 11<12；112<121 等等
		if a[i] < b[i] {
			return true
		}
		// 上面反过来的情况
		if a[i] > b[i] {
			return false
		}
		// 11和12，i=0的时候，第一位相等，就会走到这，比较下一位
		i++
	}
	if l1 == l2 {
		// 123和123 即a串等于b串
		return true
	}
	// 123 和 12312
	// a串比b串短，而且前面的部分相等
	if i == l1 {
		// 这个时候比较 a=123 和 b的后两位12
		// 123和12的前两位也相同，继续比较a的子串3和之前b的子串12
		// 因为3>12 所以 a>b
		// 这里是一个递归操作
		return less(a, b[i:])
	}
	// a串比b串长，而且前面的部分相同的情况
	// 比如上面 a=123和b的子串12(新的b串)比的时候就会走到这里来
	return less(a[i:], b)
}

// 实现函数
func minPatch(arr []int) (res string) {
	// 整型转字符串
	arrString := intArrayToStringArray(arr)
	// 排序
	sort.Sort(arrString)
	// 拼接
	for i := range arrString {
		res += arrString[i]
	}
	return
}

func TestMinPatch(t *testing.T) {
	arrInt := []int{3, 32, 321}
	res := minPatch(arrInt)
	fmt.Printf("%s", res)
}

/*
题意:
给定一个正整数的数组，[3, 32, 321, 324]，把这些数拼起来，要求结果最小
思路:
本质上是一个排序题，我们要把"最小"的数排在前面，但这个最小，不是数学上的数值最小，也不是字典序的最小，而是基于题意的一种特殊的"最小"。
比如，
按数值大小拼的话，3和32，3<32 但是332>323，不行
按字典序大小拼的话，3和321，32<321，但是32321>32132，不行
具体的less规则在代码中有详细注释
总结:
这题写的还是很有意思的，尽管排序的部分没有自己写，但是这里排序反而没那么重要，
确定了less规则之后，用快排序和堆排序都可以，甚至简单的用冒泡也行，精髓还是确定less规则
*/
