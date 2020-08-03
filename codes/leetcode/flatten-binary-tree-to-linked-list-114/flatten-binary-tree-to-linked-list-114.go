/**
 * 给定一个二叉树，原地将它展开为一个单链表。
 * 注意需要按前序顺序
 */
package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			flatten(root.Left)
		}
		if root.Right != nil {
			flatten(root.Right)
		}
		tmpOrigRight := root.Right
		root.Right = root.Left
		root.Left = nil

		tmp := root
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = tmpOrigRight
	} else {
		// do nothing
	}
}

func main() {
	// do nothing
}
