package leetcode

import (
	"container/heap"
	"fmt"
	"testing"
)

/*

Given an array A of integers, we must modify the array in the following way: we choose an i and replace A[i] with -A[i],
and we repeat this process K times in total.  (We may choose the same index i multiple times.)
Return the largest possible sum of the array after modifying it in this way.

Example 1:

Input: A = [4,2,3], K = 1
Output: 5
Explanation: Choose indices (1,) and A becomes [4,-2,3].
Example 2:

Input: A = [3,-1,0,2], K = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and A becomes [3,1,0,2].
Example 3:

Input: A = [2,-3,-1,5,-4], K = 2
Output: 13
Explanation: Choose indices (1, 4) and A becomes [2,3,-1,5,4].

Note:

1 <= A.length <= 10000
1 <= K <= 10000
-100 <= A[i] <= 100

 */

type myHeap []int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func largestSumAfterKNegations(A []int, K int) int {
	h := myHeap(A)
	heap.Init(&h)
	for i:=0; i<K; i++ {
		a := heap.Pop(&h)
		b, ok := a.(int)
		if !ok {
			break
		}
		b = -b
		heap.Push(&h, b)
	}
	sum := 0
	for _, a := range h {
		sum += a
	}
	return sum
}

func TestInit0(t *testing.T) {
	res := largestSumAfterKNegations([]int{4,2,3}, 1)
	fmt.Println(res)
}
