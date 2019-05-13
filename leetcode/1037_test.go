package leetcode

import (
	"fmt"
	"testing"
)

func isBoomerang(points [][]int) bool {
	return helper(points[0], points[1], points[2])
}

func helper(a, b, c []int) bool {
	p1, p2 := b[0] - a[0], c[0] - b[0]
	m1, m2 := b[1] - a[1], c[1] - b[1]
	if p1 * m2 == p2 * m1 {
		return false
	}
	return true
}

func TestIsBoomRange(t *testing.T) {
	a := [][]int{
		{1,1},
		{1,1},
		{3,2},
	}
	fmt.Println(isBoomerang(a))
}