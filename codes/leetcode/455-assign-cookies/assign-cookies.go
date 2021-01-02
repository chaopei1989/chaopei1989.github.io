package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	childIndex := 0
	for _, cookie := range s {
		if childIndex < len(g) {
			if cookie >= g[childIndex] {
				ret++
				childIndex++
			}
		}
	}
	return ret
}

func main() {
	findContentChildren([]int{10,9,8,7}, []int{5,6,7,8})
}