/*337. 打家劫舍 III
题目链接: https://leetcode-cn.com/problems/house-robber-iii/

在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:

输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.
示例 2:

输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.
*/
package main

import (
	"fmt"
)

// TreeNode represents a tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var mapRobWith = make(map[*TreeNode]int)

var mapRobWithout = make(map[*TreeNode]int)

// 动态规划做法, 最优子结构是:
func rob(root *TreeNode) int {
	return max(robWith(root), robWithout(root))
}

func robWith(root *TreeNode) int {
	var res int
	if root == nil {
		res = 0
	} else {
		fmt.Printf("robWith: %d\n", root.Val)
		var ok bool
		if res, ok = mapRobWith[root]; !ok {
			res = root.Val + robWithout(root.Left) + robWithout(root.Right)
			mapRobWith[root] = res
		}
	}
	return res
}

func robWithout(root *TreeNode) int {
	var res int
	if root == nil {
		res = 0
	} else {
		fmt.Printf("robWithout: %d\n", root.Val)
		var ok bool
		if res, ok = mapRobWithout[root]; !ok {
			res = max(robWith(root.Left), robWithout(root.Left)) +
				max(robWith(root.Right), robWithout(root.Right))
			mapRobWithout[root] = res
		}
	}
	return res
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func main() {
	var r *TreeNode
	r = new(TreeNode)
	r.Val = 3

	var cl, cr *TreeNode
	cl = new(TreeNode)
	cr = new(TreeNode)
	cl.Val = 2
	cr.Val = 3

	var cc1, cc2 *TreeNode
	cc1 = new(TreeNode)
	cc2 = new(TreeNode)
	cc1.Val = 3
	cc2.Val = 1

	r.Left = cl
	r.Right = cr
	cl.Right = cc1
	cr.Right = cc2

	ret := rob(r)
	fmt.Println(ret)
}
