package jianzhioffer

import (
	"fmt"
	"testing"
)

func ArrangeString(str string) {
	var dic = make(map[byte]int)

	for i := range str {
		dic[str[i]]++
	}
	printString("", dic)
}

func printString(str string, dic map[byte]int) {
	if len(dic) == 0 {
		fmt.Println(str)
		return
	}
	for ch := range dic {
		// map是传引用 每次要传一个新的map
		newDic := make(map[byte]int)
		for cc, v := range dic {
			if cc == ch {
				if v-1 > 0 {
					newDic[cc] = v - 1
				}
				continue
			}
			newDic[cc] = v
		}
		printString(str+string(ch), newDic)
	}
}

func TestArrangeString(t *testing.T) {
	var str string
	//
	str = "ab"
	fmt.Printf("***字符串:%v***\n", str)
	ArrangeString(str)
	//
	str = "abc"
	fmt.Printf("***字符串:%v***\n", str)
	ArrangeString(str)
	//
	str = "abbc"
	fmt.Printf("***字符串:%v***\n", str)
	ArrangeString(str)
	//
	str = "hello"
	fmt.Printf("***字符串:%v***\n", str)
	ArrangeString(str)
}

/*
题意:
给定一个字符串，输出所有排列。比如 ab 可以输出 ab ba；abb 可以输出 abb bab bba；abc可以输出 abc acb bac bca cab cba
思路:
先选第一个字母，然后拿着剩下的字母去生成一个子串，加起来就是一种情况。
拿有重复的 abb 来举例
可以得到起始的字母字典 [a:1, b:2]
f([a:1, b:2]) = ["a"+f([b:2]), "b"+f([a:1, b:1])]
*/
