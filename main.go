package main

import "fmt"

func main() {
	var n int
	n = 3
	fmt.Println(n)
	var routes [3][2]int
	routes[0] = [2]int{1, 2}
	routes[1] = [2]int{2, 3}
	routes[2] = [2]int{3, 1}
	fmt.Println(routes)
	record := make(map[int]map[int]bool)
	for i := 1; i <= n; i++ {
		tmp := make(map[int]bool)
		record[i] = tmp
	}

	for _, v := range routes {
		i := v[0]
		record[i][v[1]] = true
	}
	fmt.Println(record)
	return
}

func df(dp [][]int, i, j int) {
	for k := 0; k < len(dp); k++ {
		if dp[j][k] == 1 && j != k {
			dp[i][k] = 1
			if k > j {
				df(dp, j, k)
			}
		}
	}
}

func test() {
	var n int
	//fmt.Scan(&n, &m)
	n = 3
	//m = 3
	var routes [3][2]int
	//for i:=0; i<m; i++ {
	//	var tmp[2]int
	//	fmt.Scan(&tmp[0], &tmp[1])
	//	routes = append(routes, tmp)
	//}
	routes[0] = [2]int{1, 2}
	routes[1] = [2]int{2, 3}
	routes[2] = [2]int{3, 1}
	var dp [][]int
	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}
	for _, v := range routes {
		i, j := v[0]-1, v[1]-1
		dp[i][j] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if dp[i][j] == 1 {
				df(dp, i, j)
			}
		}

	}
	fmt.Println(dp)
	in := make([]int, n)
	out := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 1 {
				in[i]++
				out[j]++
			}
		}
	}
	fmt.Println(in, out)
	res := 0
	for i := 0; i < n; i++ {
		if in[i] > out[i] {
			res++
		}
	}
	fmt.Println(res)
}
