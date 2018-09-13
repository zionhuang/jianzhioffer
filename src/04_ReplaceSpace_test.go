package src

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	fmt.Println(replace("hello world"))
	fmt.Println(replace("hi huang how are you"))
}

func replace(str string) (res string) {
	numOfSpace := 0
	for _, v := range str {
		if v == ' ' {
			numOfSpace++
		}
	}
	newLen := len(str) + 2*numOfSpace
	byteArr := make([]byte, newLen)
	i := len(str) - 1
	j := newLen - 1
	for i >= 0 {
		if str[i] == ' ' {
			byteArr[j-2], byteArr[j-1], byteArr[j] = '%', '2', '0'
			j -= 3
		} else {
			byteArr[j] = str[i]
			j--
		}
		i--
	}
	return string(byteArr)
}
