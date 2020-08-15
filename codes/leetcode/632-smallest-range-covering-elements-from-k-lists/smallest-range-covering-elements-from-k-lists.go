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

// TreeNode present 小根堆
type TreeNode struct {
	Val      int
	ArrIndex int
	Left     *TreeNode
	Right    *TreeNode
}

func (root *TreeNode, last *TreeNode) Add(val int, arrIndex int) {

}

func (root *TreeNode) popTop(val int, arrIndex int) {

}

func smallestRange(nums [][]int) []int {
	// invalid integer
	var min int
	var max int
	var minArrIndex int
	var maxArrIndex int
	// 记录各个列表的当前游标
	var cursors []int = make([]int, len(nums))

	// 1. 各个列表各挑一个, 初始化
	for i := 0; i < len(nums); i++ {
		// 游标初始化
		cursors[i] = 0
		// 最大最小值初始化
		if i == 0 {
			min = nums[0][0]
			max = min
			minArrIndex = 0
			maxArrIndex = 0
		} else {
			if nums[i][0] < min {
				min = nums[i][0]
				minArrIndex = i
			}
			if nums[i][0] > max {
				max = nums[i][0]
				maxArrIndex = i
			}
		}
	}

	// 2. 每次从最小的

}
