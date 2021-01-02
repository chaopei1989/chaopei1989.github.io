package main

import "container/heap"

type Item struct {
	val   int
	point []int
}

func NewItem(val int, point []int) *Item {
	res := new(Item)
	res.point = point
	res.val = val
	return res
}

// IHeap 实际是个数组
type IHeap []*Item

func (h *IHeap) Len() int {
	return len(*h)
}

func (h *IHeap) Less(i, j int) bool {
	return (*h)[i].val > (*h)[j].val
}

func (h *IHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func euclidean2(point []int) int {
	res := 0
	for _, v := range point {
		res += v * v
	}
	return res
}

func kClosest(points [][]int, K int) [][]int {
	// 构建大根堆
	smlHeap := make(IHeap, 0)
	heap.Init(&smlHeap)
	for _, point := range points {
		pointV := euclidean2(point)
		if len(smlHeap) >= K {
			if pointV < smlHeap[0].val {
				heap.Pop(&smlHeap)
				heap.Push(&smlHeap, NewItem(pointV, point))
			}
		} else {
			heap.Push(&smlHeap, NewItem(pointV, point))
		}
	}
	res := make([][]int, 0)
	for i := 0; i < K; i++ {
		var smlItem *Item = heap.Pop(&smlHeap).(*Item)
		res = append(res, smlItem.point)
	}
	return res
}
