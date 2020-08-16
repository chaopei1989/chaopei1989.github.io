/*110. 平衡二叉树
题目链接: https://leetcode-cn.com/problems/balanced-binary-tree/

给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

提示: 这里有个思维定式陷阱, 平衡二叉树并不表示最小深度和最大深度相差小于等于1, 而是左右子树的深度相差小于等于1
*/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(l, r int) int {
	if l < r {
		return r
	} else {
		return l
	}
}

func isBalancedRec(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	lbool, ldeep := isBalancedRec(root.Left)
	rbool, rdeep := isBalancedRec(root.Right)
	delt := ldeep - rdeep
	return (lbool && rbool && delt <= 1 && delt >= -1), max(ldeep, rdeep) + 1
}

func isBalanced(root *TreeNode) bool {
	res, _ := isBalancedRec(root)
	return res
}
