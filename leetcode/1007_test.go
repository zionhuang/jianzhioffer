package leetcode

func minDominoRotations(A []int, B []int) int {
	flag := []int{-2, -2, -2, -2, -2, -2}
Label1:
	for i := 0; i < 6; i++ {
		var a, b int
		for j := 0; j < len(A); j++ {
			if A[j] != i+1 && B[j] != i+1 {
				continue Label1
			}
			if A[j] != i+1 {
				a++
			}
			if B[j] != i+1 {
				b++
			}
		}
		if a < b {
			flag[i] = a
		} else {
			flag[i] = b
		}
	}
	res := -1
	for _, m := range flag {
		if m == -2 {
			continue
		}
		if res == -1 {
			res = m
			continue
		}
		if m < res {
			res = m
		}
	}
	return -1
}
