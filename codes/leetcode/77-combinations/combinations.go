package main

import "fmt"
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	combineDFS(n, k, 1, nil, &res)
	return res
}

func combineDFS(n int, k int, s int, curr []int, res *[][]int) {
	if len(curr) == k {
		*res = append(*res, curr)
		return
	}
	if s > n {
		return
	}
	for i:=s; i<=n; i++ {
		newCurr := make([]int, len(curr)+1)
		copy(newCurr, curr)
		newCurr[len(curr)] = i
		combineDFS(n, k, i+1, newCurr, res)
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
