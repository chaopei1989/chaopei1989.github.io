package main

import (
	"fmt"
	"strconv"
)

type Result struct {
	nums  []int
	last  int
	first int
}

func (res *Result) Next(index int) bool {
	if res.first != -1 && index != res.last+1 {
		panic("index not seq")
	}
	if res.first < 0 {
		res.first = index
		res.last = index
		return true
	} else {
		if res.nums[index] == res.nums[res.last]+1 {
			res.last = index
			return true
		} else {
			return false
		}
	}
}

func (res *Result) Get() string {
	if res.last == res.first {
		return strconv.Itoa(res.nums[res.first])
	} else {
		return fmt.Sprintf("%d->%d", res.nums[res.first], res.nums[res.last])
	}
}

func NewResult(nums []int) *Result {
	ret := new(Result)
	ret.nums = nums
	ret.last = -1
	ret.first = -1
	return ret
}

func summaryRanges(nums []int) []string {
	ret := make([]string, 0)
	var tmp *Result
	for i := 0; i < len(nums); i++ {
		if tmp == nil {
			tmp = NewResult(nums)
			// must success
			tmp.Next(i)
		} else {
			if r := tmp.Next(i); !r {
				ret = append(ret, tmp.Get())
				i--
				tmp = nil
			}
		}
	}
	if tmp != nil {
		ret = append(ret, tmp.Get())
	}
	return ret
}

func main() {
	fmt.Println(summaryRanges([]int{
		0, 1, 2, 4, 5, 7,
	}))
}
