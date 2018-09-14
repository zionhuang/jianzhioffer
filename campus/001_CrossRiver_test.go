package campus

import (
	"fmt"
	"math"
	"testing"
)

func TestAa(t *testing.T) {
	arr := [][]int{
		{1, 0, 13, 34},
		{8, 9, 10, 12},
	}
	res := river(arr)
	fmt.Println(res)
}

func river(arr [][]int) int {
	r := len(arr)
	if r == 0 {
		return 0
	}
	c := len(arr[0])
	var dp [][]int
	for i := 0; i < r; i++ {
		tmp := make([]int, c)
		dp = append(dp, tmp)
	}
	for i := 1; i < r; i++ {
		for j := 0; j < c; j++ {
			min := math.MaxInt64
			if j >= 1 && arr[i-1][j-1] < min {
				min = arr[i-1][j-1]
			}
			if arr[i-1][j] < min {
				min = arr[i-1][j]
			}
			if j < c-1 && arr[i-1][j+1] < min {
				min = arr[i-1][j+1]
			}
			dp[i][j] = min + arr[i][j]
		}
	}
	min := math.MaxInt64
	for _, v := range arr[r-1] {
		if v < min {
			min = v
		}
	}
	return min
}
