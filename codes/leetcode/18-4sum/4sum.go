package main

import "sort"

type NumArray []int

func (h *NumArray) Len() int {
	return len(*h)
}

func (h *NumArray) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *NumArray) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// fourSum 入口函数
func fourSum(nums []int, target int) [][]int {
	res := make([][]int, 0)
	na := NumArray(nums)

	// 1. 保证数组有序性(从小到大)
	sort.Sort(&na)

	// 2. 四指针, 前面两个指针(a,b)控制两层循环, 后两个指针(c,d)及时调整
	for a := 0; a < len(na); a++ {
		if a > 0 && na[a] == na[a-1] {
			// 重复
			continue
		}
		for b := a + 1; b < len(na); b++ {
			if b > a+1 && na[b] == na[b-1] {
				// 重复
				continue
			}

			// 后两指针
			c := b + 1
			d := len(na) - 1
			for c < d {
				v := na[a] + na[b] + na[c] + na[d]
				if v > target {
					for d-1 > c && na[d] == na[d-1] {
						d--
					}
					d--
				} else if v < target {
					for c+1 < d && na[c] == na[c+1] {
						c++
					}
					c++
				} else {
					res = append(res, []int{na[a], na[b], na[c], na[d]})
					for d-1 > c && na[d] == na[d-1] {
						d--
					}
					d--
					for c+1 < d && na[c] == na[c+1] {
						c++
					}
					c++
				}
			}
		}
	}

	return res
}

func main() {

}
