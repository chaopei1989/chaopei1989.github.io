package main

func copyArrNewItem(orig []int, item int) []int {
	newArr := make([]int, len(orig)+1)
	copy(newArr, orig)
	newArr[len(orig)] = item
	return newArr
}

func combinationSum3DFS(s int, c int, arr []int, k int, n int, res *[][]int) {
	for i := s; i <= 9; i++ {
		newArr := copyArrNewItem(arr, i)
		newC := c + i
		if newC > n || len(newArr) > k {
			continue
		} else if newC < n {
			if len(newArr) < k {
				combinationSum3DFS(i+1, newC, newArr, k, n, res)
			} else {
				// == k
				continue
			}
		} else {
			if len(newArr) == k {
				*res = append(*res, newArr)
			} else {
				continue
			}
		}
	}
}

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	combinationSum3DFS(1, 0, nil, k, n, &res)
	return res
}
