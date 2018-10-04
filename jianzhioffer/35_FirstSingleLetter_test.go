package jianzhioffer

import (
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	str := "helloheo"
	res := firstSingleLetter(str)
	fmt.Printf("%v", string(res))
}

func firstSingleLetter(str string) byte {
	var dic map[byte]int
	dic = make(map[byte]int)
	for i := range str {
		dic[str[i]]++
	}
	for i := range str {
		if dic[str[i]] == 1 {
			return str[i]
		}
	}
	return 0
}

/*
题意:
给定一个字符串，如："abcdefg"，找到第一个单字符
思路:
第一次遍历字符串，用一个字典把各个字符出现的次数记录下来
第二次遍历，把第一个只出现一次的字符输出
*/
