package main

import "fmt"

// 未完成...
// 1. array copy 的问题
// 2. dfs 模板总结的问题
// 3. hash 的问题
// 4. 判重问题
func findSubsequencesRec(nums []int, cursor int, curr []int, res [][]int) [][]int {
	if len(curr) > 1 {
		res = append(res, curr)
	}
	arrMap := make(map[int]int)
	for i := cursor + 1; i < len(nums); i++ {
		newCurr := make([]int, len(curr))
		copy(newCurr, curr)
		if nums[i] >= nums[cursor] {
			if _, ok := arrMap[nums[i]]; i > cursor+1 && ok {
				continue
			}
			arrMap[nums[i]] = 1
			res = findSubsequencesRec(nums, i, append(newCurr, nums[i]), res)
		}
	}
	return res
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	arrMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := arrMap[nums[i]];ok {
			continue
		}
		arrMap[nums[i]] = 1
		tmp := make([]int, 0)
		tmp = append(tmp, nums[i])
		res = findSubsequencesRec(nums, i, tmp, res)
	}
	return res
}

func main() {
	//res := findSubsequences([]int{1, 1, 1, 1, 1})
	//res := findSubsequences([]int{4,6,7,7,1,1})
	res := findSubsequences([]int{100,90,80,70,60,50,60,70,80,90,100})
	fmt.Println(res)
}
