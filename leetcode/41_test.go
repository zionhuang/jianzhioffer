package leetcode

import (
	"fmt"
	"testing"
)

func firstMissingPositive(nums []int) int {
	l := len(nums)
	i := 0
	for i < l {
		if nums[i] == i+1 {
			i++
			continue
		}
		if nums[i]<1 || nums[i]>l {
			nums[i] = -1
			i++
			continue
		}
		// 此时 i处的数大于0 小于等于数组长度
		if nums[nums[i]-1] == nums[i] {
			// 这是防止可能出现重复有效数字
			nums[i] = -1
			i++
			continue
		}
		nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
	}
	fmt.Println(nums)
	for i, v := range nums {
		if v != i+1 {
			return i+1
		}
	}
	return l+1
}

func TestFirstMissingPositive(t *testing.T) {
	nums := []int{7,8,9,11,12}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}



func missingNumber(nums []int) int {
	l := len(nums)
	i := 0
	var hasN bool
	for i < l {
		if nums[i] == i || nums[i] == -1 {
			i++
			continue
		}
		if nums[i] == l {
			nums[i] = -1
			hasN = true
			i++
			continue
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
	if hasN == false {
		return l
	}
	for i := range nums {
		if nums[i] == -1 {
			return i
		}
	}
	return l
}

func TestMissingNumber(t *testing.T) {
	nums := []int{3, 1, 0}
	res := missingNumber(nums)
	fmt.Println(res)
}