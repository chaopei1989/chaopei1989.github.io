/*109. 有序链表转换二叉搜索树
题目链接: https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
*/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	// 1. 先查总节点数, 构建初步的Tree
	var root, last *TreeNode
	var count int = 0
	for tmp := head; tmp != nil; tmp = tmp.Next {
		count++
		node := new(TreeNode)
		node.Val = tmp.Val
		if root == nil {
			root = node
		}
		if last != nil {
			last.Right = node
		}
		last = node
	}
	// 2. 折半
	return halfTree(root, nil, false, count)
}

func halfTree(head *TreeNode, parent *TreeNode, left bool, count int) *TreeNode {
	// 如果总共就2个节点了, 直接返回
	if count <= 2 {
		return head
	}
	dst := count / 2
	var tmp, tmpParent *TreeNode
	tmp = head
	for i := 0; i < dst; i++ {
		tmpParent = tmp
		tmp = tmp.Right
	}
	if parent != nil {
		if left {
			parent.Left = tmp
		} else {
			parent.Right = tmp
		}
	}
	tmp.Left = head
	tmpParent.Right = nil
	halfTree(tmp.Left, tmp, true, dst)
	halfTree(tmp.Right, tmp, false, count-dst-1)
	return tmp
}
