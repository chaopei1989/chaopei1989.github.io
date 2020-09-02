package main

import "fmt"

func A(nums []int) ([]int, int, int) {
	if len(nums) == 1 {
		return []int{}, nums[0], 0
	}
	// 去掉左边的
	rest, blRes := B(nums[1:])
	var lRes, res, rRes int
	res = 0
	if len(rest) > 0 {
		_, res, _ = A(rest)
	}
	lRes = nums[0] + res

	// 去掉右边的
	rest, brRes := B(nums[0 : len(nums)-1])
	res = 0
	if len(rest) > 0 {
		_, res, _ = A(rest)
	}
	rRes = nums[len(nums)-1] + res
	if lRes >= rRes {
		return nums[1:], lRes, blRes
	} else {
		return nums[0 : len(nums)-1], rRes, brRes
	}
}

func B(nums []int) ([]int, int) {
	if len(nums) == 1 {
		return []int{}, nums[0]
	}
	// 去掉左边的
	rest, _, _ := A(nums[1:])
	var lRes, res, rRes int
	res = 0
	if len(rest) > 0 {
		_, res = B(rest)
	}
	lRes = nums[0] + res

	// 去掉右边的
	rest, _, _ = A(nums[0 : len(nums)-1])
	res = 0
	if len(rest) > 0 {
		_, res = B(rest)
	}
	rRes = nums[len(nums)-1] + res
	if lRes >= rRes {
		return nums[1:], lRes
	} else {
		return nums[0 : len(nums)-1], rRes
	}
}

func PredictTheWinner(nums []int) bool {
	_, a, b := A(nums)
	return a >= b
}

func main() {
	fmt.Println(PredictTheWinner([]int{10,17,11,16,17,9,14,17,18,13,11,4,17,18,15,3,13,10,6,8}))
}
