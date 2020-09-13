package main

import "sort"

func copyArrNewItem(orig []int, item int) []int {
	newArr := make([]int, len(orig)+1)
	copy(newArr, orig)
	newArr[len(orig)] = item
	return newArr
}

func isEqual(left []int, right []int) bool {
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}
	return true
}

func isIn(queue [][]int, item []int) bool {
	for i := 0; i < len(queue); i++ {
		if isEqual(queue[i], item) {
			return true
		}
	}
	return false
}

func combinationSum2DFS(s int, c int, arr []int, candidates []int, target int, res *[][]int)  {
	for i := s; i < len(candidates); i++ {
		newArr := copyArrNewItem(arr, candidates[i])
		newC := c + candidates[i]
		if newC > target {
			continue
		} else if newC < target {
			combinationSum2DFS(i+1, newC, newArr, candidates, target, res)
		} else {
			if len(*res) == 0 || !isIn(*res, newArr) {
				*res = append(*res, newArr)
			} else {
				continue
			}
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	combinationSum2DFS(0, 0, nil, candidates, target, &res)
	return res
}
