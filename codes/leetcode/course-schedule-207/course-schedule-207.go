// 207. 课程表
// 题目链接: https://leetcode-cn.com/problems/course-schedule/
//
// 你这个学期必须选修 numCourse 门课程，记为 0 到 numCourse-1 。
//
// 在选修某些课程之前需要一些先修课程。 例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示他们：[0,1]
//
// 给定课程总量以及它们的先决条件，请你判断是否可能完成所有课程的学习？
//
// 示例 1:
//
// 输入: 2, [[1,0]]
// 输出: true
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要完成课程 0。所以这是可能的。
//
// 示例 2:
//
// 输入: 2, [[1,0],[0,1]]
// 输出: false
// 解释: 总共有 2 门课程。学习课程 1 之前，你需要先完成​课程 0；并且学习课程 0 之前，你还应先完成课程 1。这是不可能的。
//
// 分析:
//
// 该题本质是要判断一个图是否为 DAG.
package main

import (
	"fmt"
)

// 使用邻接表(adjacency), 拓扑排序实现
func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}
	// 1. 构建邻接表和入度表
	var adjacency [][]int
	var indegrees []int
	// 以出为索引的邻接表单项
	adjacency = make([][]int, numCourses)
	// 以入为索引的入度表
	indegrees = make([]int, numCourses)
	for _, edge := range prerequisites {
		var from, to int
		from = edge[1] // 出
		to = edge[0]   // 入

		var adj []int
		adj = adjacency[from]
		if adj == nil {
			adj = make([]int, 0)
			adjacency[from] = adj
		}
		adjacency[from] = append(adj, to)
		indegrees[to]++
	}

	// 2. 找到入度为 0 的元素, 广度遍历
	// 先构建临时数组来存放入度为 0 的元素
	var zeroIndegreeIndex []int
	zeroIndegreeIndex = make([]int, 0)
	for i, indegree := range indegrees {
		if indegree == 0 {
			zeroIndegreeIndex = append(zeroIndegreeIndex, i)
		}
	}
	var numTmp int
	numTmp = numCourses
	for {
		if len(zeroIndegreeIndex) > 0 {
			// 入度为 0 的元素的邻接元素入度-1
			var zii int
			zii = zeroIndegreeIndex[0]
			if adjacency[zii] != nil {
				for _, to := range adjacency[zii] {
					indegrees[to]--
					if indegrees[to] == 0 {
						zeroIndegreeIndex = append(zeroIndegreeIndex, to)
					}
				}
			}
			zeroIndegreeIndex = zeroIndegreeIndex[1:]
			numTmp--
		} else {
			break
		}
	}

	return numTmp == 0
}

func canFinish1(numCourses int, prerequisites [][]int) bool {
	return false
}

func main() {
	ret := canFinish(3, [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{2, 1},
	})
	fmt.Printf("main: %t\n", ret)
	// tmp := []int{1, 2, 3}
	// i := 4
	// for _, t := range tmp {
	// 	fmt.Printf("main: %d\n", t)
	// 	tmp = append(tmp, i)
	// 	i++
	// }
}
