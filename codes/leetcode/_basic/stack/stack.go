package main

import (
	"fmt"
)

/***************************************************************
 *
 * 使用 slice 来实现栈结构, 效率高, 但由于 slice 中数组共享机制, 有内存问题.
 *
***************************************************************/

type StackArray []interface{}

func (self *StackArray) Push(item interface{}) {
	*self = append(*self, item)
}

func (self *StackArray) Pop() interface{} {
	ret := (*self)[len((*self))-1]
	*self = (*self)[0 : len((*self))-1]
	return ret
}

func (self *StackArray) Len() int {
	return len(*self)
}

/***************************************************************
 *
 * 		下面是测试程序
 *
***************************************************************/

func main() {
	stackArray := new(StackArray)
	a, b, c, d, e := 1, 2, 3, 4, 5
	stackArray.Push(a)
	stackArray.Push(b)
	stackArray.Push(c)
	stackArray.Push(d)
	stackArray.Push(e)
	fmt.Println(stackArray.Pop().(int))
	fmt.Println(stackArray.Pop().(int))
	fmt.Println(stackArray.Pop().(int))
	fmt.Println(stackArray.Pop().(int))
	fmt.Println(stackArray.Pop().(int))

	fmt.Printf("len=%d\n", stackArray.Len())
}
