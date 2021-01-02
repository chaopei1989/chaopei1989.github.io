package main

import "sort"

type P [][]int

func (self P) Len() int {
	return len(self)
}

func (self P) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self P) Less(i, j int) bool {
	// 首位降序
	if self[i][0] > self[j][0] {
		return true
	} else if self[i][0] < self[j][0] {
		return false
	} else {
		// 次位升序
		return self[i][1] < self[j][1]
	}
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(P(people))
	res := make([][]int, 0)
	for i := 0; i < len(people); i++ {
		if people[i][1] >= len(res) {
			res = append(res, people[i])
		} else {
			person := people[i]
			idx := person[1]
			res = append(res[:idx], append([][]int{person}, res[idx:]...)...)
			//res = append(append(res[0:people[i][1]], people[i]), res[people[i][1]:]...)
		}
	}
	return res
}

