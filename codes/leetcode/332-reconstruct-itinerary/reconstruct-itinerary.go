package main

import "sort"

var (
	ticketsMap = make(map[string][]string)
)
const (
	FROM = "JFK"
)

func findItinerary(tickets [][]string) []string {
	// 1. 所有机票以 [from, [to]] 的形式存储, [to] 需要按字典序从小到大排序
	for _, t := range tickets {
		arr := ticketsMap[t[0]]
		if arr == nil {
			arr = make([]string, 0)
		}
		arr = append(arr, t[1])
		ticketsMap[t[0]] = arr
	}
	for key := range ticketsMap {
		sort.Strings(ticketsMap[key])
	}
	// 2. Hierholzer 算法
	res := findItineraryCursive(FROM, nil)
	for i := 0; i < len(res) / 2; i++ {
		res[i], res[len(res)-i-1]= res[len(res)-i-1], res[i]
	}
	return res
}

func findItineraryCursive(from string, res []string) []string {
	if res == nil {
		res = make([]string, 0)
	}
	for {
		toArr, ok := ticketsMap[from]
		if !ok || len(toArr) == 0 {
			break
		}
		to := toArr[0]
		ticketsMap[from] = toArr[1:]
		res = findItineraryCursive(to, res)
	}
	// 倒序进数组
	res = append(res, from)
	return res
}

func main()  {
	
}
