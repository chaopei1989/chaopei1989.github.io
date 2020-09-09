package main

import "fmt"

func combinationSumDFS(candidates []int, target int, currRes, s int, arr []int, res *[][]int) {
	for i := s; i < len(candidates); i++ {
		newArr := make([]int, len(arr) + 1)
		copy(newArr, arr)
		newArr[len(arr)] = candidates[i]
		newRes := currRes + candidates[i]
		if newRes > target {
			continue
		} else if newRes < target {
			combinationSumDFS(candidates, target, newRes, i, newArr, res)
		} else {
			*res = append(*res, newArr)
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	combinationSumDFS(candidates, target, 0, 0, nil, &res)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{8,7,4,3}, 11))
}
