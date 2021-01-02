package main

import "fmt"

type IncStack struct {
	origin     []int
	indexQueue []int
	windowSize int
	cursor     int
}

func NewIncStack(origin []int, windowSize int) *IncStack {
	ret := new(IncStack)
	ret.origin = origin
	ret.indexQueue = make([]int, 0)
	ret.windowSize = windowSize
	return ret
}

func (stack *IncStack) push(index int) {
	stack.cursor++
	for i := len(stack.indexQueue) - 1; i >= -1; i-- {
		// 找到了插入点
		if i == -1 || stack.get(stack.indexQueue[i]) >= stack.get(index) {
			dst := stack.indexQueue[0 : i+1]
			dst = append(dst, index)
			// 超size
			if dst[0] + 1 <= stack.cursor - stack.windowSize {
				dst = dst[1:]
			}
			stack.indexQueue = dst
			return
		}
	}
}

func (stack *IncStack) get(index int) int {
	return stack.origin[index]
}

func (stack *IncStack) max() int {
	return stack.get(stack.indexQueue[0])
}

func maxSlidingWindow(nums []int, k int) []int {
	stack := NewIncStack(nums, k)
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		stack.push(i)
		if stack.cursor >= k {
			ret = append(ret, stack.max())
		}
	}
	return ret
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
}
