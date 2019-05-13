package campus

import (
	"fmt"
	"strconv"
	"testing"
)

func strToNum(num string) int {
	if num == "" {
		return 0
	}
	if num[0] == '0' {
		if len(num) == 1 {
			return 0
		}
		if num[1] == 'x' {
			if len(num) == 2 {
				return 0
			}
			return sixthToTen(num[2:])
		}
		return eightToTen(num[1:])
	}
	res, _ := strconv.Atoi(num)
	return res
}

func sixthToTen(num string) int {
	var res int
	for i := range num {
		tmp, _ := strconv.Atoi(string(num[i]))
		if i == 0 {
			res = tmp
			continue
		}
		res = res * 16 + tmp
	}
	return res
}

func eightToTen(num string) int {
	var res int
	for i := range num {
		tmp, _ := strconv.Atoi(string(num[i]))
		if i == 0 {
			res = tmp
			continue
		}
		res = res * 8 + tmp
	}
	return res

}

func TestStrToNum(t *testing.T) {
	num := "010"
	res := strToNum(num)
	fmt.Println(res)
}

func TestSsss(t *testing.T) {
	num := "1+011+0x11"
	res := ssss(num)
	fmt.Println(res)
}

func ssss(input string) int {
	var res int
	var pre int
	var cur int
	var flag bool
	for _, v := range input {
		switch v {
		case '+':
			tmp := strToNum(input[pre:cur])
			if !flag {
				res += tmp
			} else {
				res -= tmp
			}
			flag = false
			pre = cur+1
			cur++
		case '-':
			tmp := strToNum(input[pre:cur])
			if !flag {
				res += tmp
			} else {
				res -= tmp
			}
			flag = true
			pre = cur+1
			cur++
		default:
			cur++
		}
	}
	tmp := strToNum(input[pre:])
	if !flag {
		res += tmp
	} else {
		res -= tmp
	}
	return res
}