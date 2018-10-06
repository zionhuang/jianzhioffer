package leetcode

import (
	"testing"
	"fmt"
)

func isHead(str, head string) bool {
	l := len(head)
	if str[:l] == head {
		return true
	}
	return false
}

func isOk(str string, dic map[string]int) bool {
	for k := range dic {
		if isHead(str, k) {
			l := len(k)
			newDic := make(map[string]int)
			for key := range dic {
				if key == k {
					if dic[key]-1 != 0 {
						newDic[key] = dic[key] - 1
					}
				} else {
					newDic[key] = dic[key]
				}
			}
			if len(newDic) == 0 {
				return true
			}
			return isOk(str[l:], newDic)
		}
	}
	return false
}



func TestIsOk(t *testing.T) {
	str := "helloworld"
	dic := make(map[string]int)
	dic["hello"] = 1
	l := 5
	res := isOk(str[:l], dic)
	fmt.Println(res)

}

func substring(str string, words []string) (res []int) {
	if str == "" || words == nil || len(words) == 0 {
		return
	}
	dic := make(map[string]int)
	cnt := 0
	wl := len(words[0])
	for i := range words {
		if wl != len(words[i]) {
			return
		}
		dic[words[i]]++
		cnt++
	}
	sl := len(str)
	wsl := cnt * wl
	for i:=0; i<=sl-wsl; i++ {
		if isOk(str[i:i+wsl], dic) {
			res = append(res, i)
		}
		fmt.Println(i)
	}
	return
}

func TestSubstring(t *testing.T) {
	str := "helloworldhello"
	words := []string{
		"hello",
	}
	res := substring(str, words)
	fmt.Println(res)
}

/*
题意:
You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once
and without any intervening characters.
给定一个字符串str，和一个“单词”数组words，数组里的“单词”长度都一样，从str中找到所有包含所有单词的子串
str = "barfoothefoobarman",
words = ["foo","bar"]
输出:
[0, 9] 因为 str[0:6](barfoo)和str[9:15](foobar)都包含foo和bar，符合条件，输出[0, 9]是两个符合条件的子串的起始下标
思路:
以一个窗口进行滑动，窗口长度是words的总长度，要判断这个串口内的子串是否符合条件
 */