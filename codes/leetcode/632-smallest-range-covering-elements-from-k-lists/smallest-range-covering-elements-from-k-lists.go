/*632. 最小区间
题目链接: https://leetcode-cn.com/problems/smallest-range-covering-elements-from-k-lists/

你有 k 个升序排列的整数列表。找到一个最小区间，使得 k 个列表中的每个列表至少有一个数包含在其中。

我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。



示例：

输入：[[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
输出：[20,24]
解释：
列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。


提示：

给定的列表可能包含重复元素，所以在这里升序表示 >= 。
1 <= k <= 3500
-105 <= 元素的值 <= 105
对于使用Java的用户，请注意传入类型已修改为List<List<Integer>>。重置代码模板后可以看到这项改动。

提示: 该问题可以转化为, 从 k 个列表中各取一个数, 使得这 k 个数中的最大值与最小值的差最小.
从各个列表中各选一个数, 其中最小值和最大值构成的区间一定符合要求. 不断缩小这个区间.

证明: 每次都右移最小值指针, 一定能缩小区间, 且在不能右移时达到最小.
*/
package main

import (
	"container/heap"
)

type Item struct {
	Val      int
	ArrIndex int
}

// IHeap 实际是个数组
type IHeap []*Item

func (h *IHeap) Len() int {
	return len(*h)
}

func (h *IHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

var (
	smlHeap IHeap
	cursors []int
	max int
)

func smallestRange(nums [][]int) []int {
	var from, to int
	var delt int = -1
	// 构建小根堆
	smlHeap = make(IHeap, 0)
	heap.Init(&smlHeap)
	// 记录各个列表的当前游标
	cursors = make([]int, len(nums))

	// 1. 各个列表各挑第一个, 进行初始化
	for i := 0; i < len(nums); i++ {
		// 游标初始化
		cursors[i] = 0
		tmp := new(Item)
		tmp.ArrIndex = i
		tmp.Val = nums[i][0]
		heap.Push(&smlHeap, tmp)
		if i == 0 {
			max = tmp.Val
		} else {
			if max < tmp.Val {
				max = tmp.Val
			}
		}
	}

	for {
		// 2. 每次取出最小的, 计算区间并更新
		var smlItem *Item = heap.Pop(&smlHeap).(*Item)
		// max - smlItem.Val == delt 不采用, 因为区间边缘一定是大的
		if delt < 0 || max - smlItem.Val < delt {
			from = smlItem.Val
			to = max
			delt = to - from
		}
		// 3. 更新下一个游标
		cursors[smlItem.ArrIndex]++
		currCursor := cursors[smlItem.ArrIndex]
		if currCursor < len(nums[smlItem.ArrIndex]) {
			tmp := new(Item)
			tmp.ArrIndex = smlItem.ArrIndex
			tmp.Val = nums[smlItem.ArrIndex][currCursor]
			heap.Push(&smlHeap, tmp)
			if tmp.Val > max {
				max = tmp.Val
			}
		} else {
			break
		}
	}

	// 3. return
	return []int{from, to}
}

func main() {

}
